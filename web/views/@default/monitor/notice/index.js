Tea.context(function () {

    this.onChangeTimeFormat = function (time) {
        return time
      };

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除所选监控消息？", function () {
            this.$post(".delete")
                .params({
                    id: id
                })
                .refresh()
        })
    }
})