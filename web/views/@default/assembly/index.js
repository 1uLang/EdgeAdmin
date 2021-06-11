Tea.context(function () {
	this.createNode = function () {
		teaweb.popup(Tea.url(".createPopup"), {
			height: "30em",
			callback: function () {
				teaweb.success("保存成功", function () {
					teaweb.reload()
				})
			}
		})
	}

	this.updateNode = function (id) {
		teaweb.popup(Tea.url(".update?NodeId="+id), {
			height: "30em",
			callback: function () {
				teaweb.success("修改成功", function () {
					teaweb.reload()
				})
			}
		})
	}

	this.deleteNode = function (nodeId) {
		let that = this
		teaweb.confirm("确定要删除这个节点吗？", function () {
			that.$post(".delete")
				.params({
					NodeId: nodeId
				})
				.refresh()
		})
	}
})