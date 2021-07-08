Tea.context(function () {

    this.onOpenCommand = function () {
        teaweb.popup(Tea.url(".install"),{
            height: "350px",
        });
    }

    this.$delay(function () {

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })
    this.onStartConfig = function (item) {
        teaweb.confirm("确定要启动吗？", function () {
            this.$post(".disport")
                .params({
                    MacCode: item.macCode,
                    Opt: 'enable',
                }).success(function (){
                window.location.reload()
            })
        })

    }

    this.onStopConfig = function (item) {
        teaweb.confirm("确定要停用吗？", function () {
            this.$post(".disport")
                .params({
                    MacCode: item.macCode,
                    Opt: 'disable',
                }).success(function (){
                window.location.reload()
            })
        })

    }

    this.onDelete = function (item) {
        teaweb.confirm("确定要卸载吗？", function () {
            this.$post(".disport")
                .params({
                    MacCode: item.macCode,
                    Opt: 'delete',
                }).success(function (){
                window.location.reload()
            })
        })
    }

    this.getStateName = function (state) {
        if(state == '1')
            return "启用中"
        else if(state == '2')
            return "已启用"
        else if(state == '3')
            return "停用中"
        else if(state == '4')
            return "已停用"
        else if(state == '5')
            return "卸载中"
        else if(state == '6')
            return "已卸载"
    }
});
  