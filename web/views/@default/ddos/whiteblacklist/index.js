Tea.context(function () {

    this.onAddNameList = function () {
        let node = this.getNodeId()
        teaweb.popup(Tea.url(".createPopup?nodeId=" + node), {
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload();
                });
            },
        });
    }

    this.getNodeId = function () {
        let node = this.nodeId

        return node
    }
    this.onDelete = function (item) {
        let node = this.getNodeId()
        teaweb.confirm("确定删除该黑白名单吗？", function () {
            this.$post(".del").params({
                Addr: [item.Address],
                Type: item.Flags,
                NodeId: node,
            }).refresh()
        })
    }

    this.showHost = function () { //重新加载该页面
        let node = ''
        if (this.nodeId === '') {    //重新加载该页面
            node = document.getElementById('selectBox').value
            this.nodeId = node
        } else {
            node = this.nodeId
        }
        window.location.href = '/ddos/whiteblacklist?nodeId=' + node
    }
    this.toShowFlags = function (flags) {
        if (flags === "blacklist") {
            return "白名单"
        } else {
            return "黑名单"
        }
    }
})