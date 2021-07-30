// LOAD JS START
$(window).load(function() {
	unblockUI();
});

function blockUI(){
	$('.loaderCenter').show();
	Liferay.Session.extend();
}

function unblockUI(){
	$('.loaderCenter').hide();
}