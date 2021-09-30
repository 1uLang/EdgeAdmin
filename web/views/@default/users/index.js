Tea.context(function () {
	this.createUser = function () {
		teaweb.popup(Tea.url(".createPopup"), {
			height: "30em",
			callback: function () {
				teaweb.success("保存成功", function () {
					teaweb.reload()
				})
			}
		})
	}

	this.deleteUser = function (userId) {
		let that = this
		teaweb.confirm("确定要删除这个用户吗？", function () {
			that.$post(".delete")
				.params({
					userId: userId
				})
				.refresh()
		})
	}

	this.onLookDetail = function(id){
		window.location = "/users/sub?parentId=" + id
	}

	this.GetSelectChan = function (event) {
		this.selectChan = event.target.value; //获取option对应的value值
		// localStorage.setItem("nfwSelectNodeId", this.selectNode);
		let chan = this.selectChan
		window.location.href = '/users?selectChan=' + chan

	}
})