package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const uploadMemoryMax = 10 * 1024 * 1024 // 2M (in bytes)
const uploadDirectory = "./uploads/"     // where to put files

func allowedFile(s string) bool {
	allowed := []string{".jpg", ".jpeg", ".mov", ".mp4",".png"}
	for _, ext := range allowed {
		if strings.HasSuffix(strings.ToLower(s), ext) {
			return true
		}
	}
	return false
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// get file handle for uploaded file
	r.ParseMultipartForm(uploadMemoryMax)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid request", 400)
		return
	}
	defer file.Close()
	if len(handler.Filename) > 0 && allowedFile(handler.Filename) {
		f, err := os.OpenFile(uploadDirectory+handler.Filename, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "invalid request", 400)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		// set modified and access time
		lastModified, err := strconv.ParseInt(r.Form.Get("lastModified"), 10, 64)
		if err == nil {
			t := time.Unix(lastModified, 0)
			os.Chtimes(uploadDirectory+handler.Filename, t, t)
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
	http.Error(w, "type not allowed", 415)
	return
}

func (as *Server) upload_video(w http.ResponseWriter, r *http.Request)() {
	os.MkdirAll(uploadDirectory, os.ModePerm)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir(uploadDirectory))))
	http.HandleFunc("/upload", handleUpload)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
