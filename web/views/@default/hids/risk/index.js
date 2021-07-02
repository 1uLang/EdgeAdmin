Tea.context(function () {
    this.seriousPer = "0%"
    this.heightPer = "0%"
    this.middlePer = "0%"
    this.lowPer = "0%"

    this.$delay(function () {
        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })

    this.onOpenDetail = function(type){
        switch(type){
            case 1:
                window.location = "/hids/risk/systemRisk"
                break
            case 2:
                window.location = "/hids/risk/weak"
                break
            case 3:
                window.location = "/hids/risk/dangerAccount"
                break
            case 4:
                window.location = "/hids/risk/configDefect"
                break
            default:
                break
        }
       
    }

    this.data = {
        seriousCount:10,
        heightCount:20,
        middleCount:10,
        lowCount:1,
        errHostCount:10,
        errWinCount:5,
        errLinuxCount:12,
        value1:{//弱口令
            allCount:206,
            errCount:15
        },
        value2:{//风险账号
            allCount:206,
            errCount:15
        },
        value3:{//缺陷配置
            allCount:206,
            errCount:15
        },
    }

});
  