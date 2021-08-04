Tea.context(function () {
	this.username = ""
	this.password = ""
	this.passwordMd5 = ""
	this.encodedFrom = window.encodeURIComponent(this.from)

	if (this.isDemo) {
		this.username = "admin"
		this.password = "123456"
	}

	this.showOTP = false

	this.isSubmitting = false

	this.$delay(function () {
		sessionStorage.setItem("leftSelectCode","dashboard")
		localStorage.removeItem("ddosSelectNodeId");
		localStorage.removeItem("nfwSelectNodeId");
		this.$find("form input[name='username']").focus()
		this.changePassword()
	});

	this.changeUsername = function () {
		this.$post("/checkOTP")
			.params({
				username: this.username
			})
			.success(function (resp) {
				this.showOTP = resp.data.requireOTP
			})
	}

	this.changePassword = function () {
		this.passwordMd5 = md5(this.password.trim());
	};

	// 更多选项
	this.moreOptionsVisible = false;
	this.showMoreOptions = function () {
		this.moreOptionsVisible = !this.moreOptionsVisible;
	};

	this.submitBefore = function () {
		this.isSubmitting = true;
	};

	this.submitDone = function () {
		this.isSubmitting = false;
	};

	this.submitSuccess = function () {
		if (this.from.length == 0) {
			window.location = "/dashboard";
		} else {
			window.location = this.from;
		}
	};

	this.submitFaild = function(res){
		let that = this
		if(res.code==100){ //如果是密码过期
			teaweb.warn(res.message,function(){
				that.onOpenResetPassword()
			})
		}else{
			teaweb.warn(res.message)
		}
	}

	this.onOpenResetPassword = function(){
		window.location = "/updatePwd"
	}
});