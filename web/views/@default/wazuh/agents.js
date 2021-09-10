Tea.context(function () {

    this.hanleOpen = function () {
        teaweb.popup(Tea.url("/hids/agents/create"), {});
    };
    this.onDelete = function (agent){

        teaweb.confirm("确定要删除所选资产吗？", function () {
            this.$post("/hids/agents/delete")
                .params({
                    agent: agent,
                }).success(function () {
                window.location.reload()
            })
        })

    }
});
