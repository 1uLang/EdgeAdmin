Tea.context(function () {

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

    this.onLookDetail = function(item){
        window.location = "/users?channelId" + item.id
    }


})