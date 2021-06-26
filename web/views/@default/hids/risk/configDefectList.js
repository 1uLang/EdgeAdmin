Tea.context(function () {
    this.nTableState = 1;

    this.onChangeTableState = function (state) {
        if (this.nTableState != state) {
            this.nTableState = state;
        }
    };

    this.getConfigDefectTypeName = function (tp) {
        if ('04011')
            return "操作系统配置缺陷";
        else if ('04012')
            return "操作系统配置缺陷";
        else if ('0402')
            return " WEB 容器配置缺陷";
        else if ('0403')
            return " 应用配置缺陷";
        else if ('0404')
            return " 数据库配置缺陷";
        else
            return "未知";
    }


    this.getStatusName = function (status) {
        switch (status) {
            case 1:
                return "已修复";
            case 2:
                return "待修复";
            case 3:
                return "修复中";
            case 4:
                return "修复失败";
            case 5:
                return "已忽略";
            default:
                "未知";
        }
    };


    this.onGoBack = function () {
        window.location = "/hids/risk/configDefect"
    };

    this.getDangerName = function (status) {
        switch (status) {
            case 1:
                return "低危";
            case 2:
                return "中危";
            case 3:
                return "高危";
            case 4:
                return "危机";
            default:
                "未评级";
        }
    };
    this.onDetail = function (item) {
        teaweb.popup(Tea.url(".configDefectDetail?macCode=" + this.macCode +
            '&riskId=' + item.riskId +
            '&state=' + this.nTableState
        ), {
            height: "30em",
        })
    };
    this.onIgnore = function (item) {
        teaweb.confirm("确定忽略该配置缺陷吗？", function () {
            this.$post(".configDefect").params({
                Opt: "ignore",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.cancelOnIgnore = function (item) {
        teaweb.confirm("确定取消忽略该配置缺陷吗？", function () {
            this.$post(".configDefect").params({
                Opt: "cancel_oignore",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.parseServerLocalIp = function (ip){
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }
})
;