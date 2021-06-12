Tea.context(function () {


    this.createPolicy = function () {
        teaweb.popup("/ddos/whiteblacklist/createPopup", {
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload()
                })
            }
        })
    }



    this.deleteGroup = function (groupId) {
        let that = this
        teaweb.confirm("确定要删除这个分组吗？", function () {
            that.$post("/servers/components/groups/delete")
                .params({
                    groupId: groupId
                })
                .success(function () {
                    teaweb.success("删除成功", function () {
                        teaweb.reload()
                    })
                })
        })
    }
})