// CARD INFO JS
$(document).ready(function() {
	sectionInfo();
});
function sectionInfo() {
	$(".section-info").click(function() {
		$(this).siblings(".info-box").slideToggle(250);
		$(this).closest(".expose").toggleClass("in");
		$('#overlay').fadeToggle(300);
	});

	$('#overlay').click(function(e) {
		$(".info-box").slideUp(250);
		$('#overlay').fadeOut(300, function() {
			$('.expose').removeClass("in");
		});
	});
}

function keepSessionAlive() {
	var milliSeconds =180000; 
	setInterval(function(){ 
		Liferay.Session.extend();
		console.log("extending session");
	}, milliSeconds);
}



function hideParent(){
	$(".modal1").css("z-index","100");	
}
function showParent(){
	$(".modal1").css("z-index","1004");
}
