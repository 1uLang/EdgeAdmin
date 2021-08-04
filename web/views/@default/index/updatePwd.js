Tea.context(function () {
	this.confirmPassword = ""
	this.password = ""
	this.passwordMd5 = ""
	this.token = ""

	this.isSubmitting = false

	this.$delay(function () {

	});

	this.changePassword = function () {
		this.passwordMd5 = md5(this.password.trim());
	};

	this.submitBefore = function () {
		this.isSubmitting = true;
	};

	this.submitDone = function () {
		this.isSubmitting = false;
	};

	this.submitSuccess = function () {
		window.location = "/index";
	};
});