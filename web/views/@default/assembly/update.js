Tea.context(function () {
	this.success = NotifySuccess("保存成功", "/assembly/update?NodeId=" + this.node.id)

	this.passwordEditing = false

	this.changePasswordEditing = function () {
		this.passwordEditing = !this.passwordEditing
	}
})