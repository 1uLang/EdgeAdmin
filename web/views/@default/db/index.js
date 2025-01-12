Tea.context(function () {
	// 添加节点
	this.createNode = function () {
		teaweb.popup("/db/createPopup", {
			height: "30em",
			callback: function () {
				teaweb.success("保存成功", function () {
					window.location.reload()
				})
			}
		})
	}

	// 修改节点
	this.updateNode = function (nodeId) {
		teaweb.popup("/db/updatePopup?nodeId=" + nodeId, {
			height: "30em",
			callback: function () {
				teaweb.success("保存成功", function () {
					window.location.reload()
				})
			}
		})
	}

	// 删除节点
	this.deleteNode = function (nodeId) {
		let that = this
		teaweb.confirm("确定要删除此数据库节点吗？", function () {
			that.$post(".delete")
				.params({
					nodeId: nodeId
				})
				.refresh()
		})
	}
	
    // 清理节点
    this.cleanNode = function (nodeId) {
        teaweb.popup("/db/cleanPopup?nodeId=" + nodeId, {
        	width: "44em",
        	height: "26em"
		})
    }

	// 显示错误信息
    this.showError = function (err) {
	    teaweb.popupTip("<span style=\"color:#db2828\">错误信息：" + err + "</span>")
    }
})