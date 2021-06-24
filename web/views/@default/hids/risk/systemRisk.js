Tea.context(function () {
    this.keyword = ""
    this.moreSelect = ""
    this.curIndex = -1
    this.showSelectValue = ""
    this.selectValue = []

    this.bShowSelectBox = false

    this.onGoBack = function () {
        window.location = "/hids/risk";
    }

    this.parseServerLocalIp = function (ip){
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }
    this.onOpenDetail = function (item) {
        window.location = "/hids/risk/systemRiskList?ip=" + item.serverIp +  '&macCode=' +item.macCode
    }
    this.dangerData = [
        {id: 1, value: "低危"},
        {id: 2, value: "中危"},
        {id: 3, value: "高危"},
        {id: 4, value: "危急"},
        {id: 5, value: "未评级"},
    ];
});