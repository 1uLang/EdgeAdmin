Tea.context(function () {

    this.Items = this.examineItems !== ""? this.examineItems.split(","): []
    this.curIndex = -1

    this.bShowCheckDetail = false
    this.pCheckDetailData = null
    this.sSelectCheckValue = ["01","02","03","04","13","14","15"]

    this.webSearchKey = ""  //网页后门
    this.searchPath = ""    //病毒木马
    this.MacCode = ""
    this.bTimeOutTip = false
    this.bShowScanPath = false

    this.dayFrom = ""
    this.dayTo = ""

    this.sTopSelectItem = [
        {id: "01", value: "系统漏洞"},
        {id: "02", value: "弱口令"},
        {id: "03", value: "风险账号"},
        {id: "04", value: "配置缺陷"},
        {id: "11", value: "病毒木马"},
        {id: "12", value: "网页后门"}
    ]
    this.sBottomSelectItem = [
        {id: "13", value: "反弹shell"},
        {id: "14", value: "异常账号"},
        {id: "15", value: "系统命令篡改"},
        {id: "16", value: "异常进程"},
        {id: "17", value: "日志异常删除"},
    ]

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })

    this.onChangeCheckState = function (state) {
        if (this.state != state) {
            this.state = state
        }
        this.refreshPage()
    }

    this.onChangeHealthNumState = function (state) {
        if (this.score != state) {
            this.score = state
        }
        this.refreshPage()
    }

    this.onChangeResultState = function (state) {
        if (this.Type != state) {
            this.Type = state
        }
        this.refreshPage()
    }
    this.refreshPage = function () {
        let url = "/hids/examine?state=" + this.state + "&score=" + this.score + "&Type=" + this.Type
        if (this.Items.length > 0) {
            url += "&examineItems=" + this.Items.toString()
        }
        window.location = url

    }

    this.parseServerLocalIp = function (ip) {
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }


    //检测是否包含元素
    this.checkSelectValue = function (index, selectValue) {
        if (selectValue) {
            for (var i = 0; i < selectValue.length; i++) {
                if (selectValue[i] == index) {
                    return true;
                }
            }
        }

        return false;
    }
    //添加/删除元素
    this.onAddSelectValue = function (index) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(index, this.Items);
        }
        if (bValue) {
            this.Items = this.Items.filter((itemIndex) => {
                return itemIndex != index;
            });
        } else {
            this.Items.push(index);
        }

        this.refreshPage()
    }


    this.getShowSelectImage = function (id) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(id, this.Items);
        }
        if (bValue) {
            return "/images/select_select.png";
        }
        return "/images/select_box.png";
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.getStatusName = function (status) {
        switch (status) {
            case 0:
                return "未体检"
            case 1:
                return "体检中"
            case 2:
                return "已完成"
            default:
                return "未知"
        }
    }

    this.getCheckNumName = function (num) {
        return (num && num > 0) ? num + "分" : "未得出"
    }

    this.getHealthName = function (score) {
        if (score =>0 && score <= 59){
            return '不健康'
        }else if (score >=60 && score <=89){
            return '亚健康'
        }else{
            return '健康'
        }
    }

    this.enters = function (index) {
        this.curIndex = index;
    }

    this.leaver = function (index) {
        this.curIndex = -1;
    }

    this.onOpenDetail = function (item) {
        window.location = "/hids/examine/detail?macCode="+item.macCode
    }

    this.onOpenCheck = function (item) {
        this.MacCode = item.macCode
        this.sSelectCheckValue = ["01","02","03","04","13","14","15"]
        this.pCheckDetailData = [
            {
                checkName: "漏洞风险检查项：",
                checkValue: [
                    {id: "01", value: "系统漏洞"},
                    {id: "02", value: "弱口令"},
                    {id: "03", value: "风险账号"},
                    {id: "04", value: "配置缺陷"},
                ]
            },
            {
                checkName: "入侵威胁检查项：",
                checkValue: this.sBottomSelectItem
            }
        ]
        if (this.pCheckDetailData) {
            this.bShowCheckDetail = true
        } else {
            this.bShowCheckDetail = false
        }
    }

    this.onSelectCheckValue = function (index) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(index, this.sSelectCheckValue);
        }
        if (bValue) {
            this.sSelectCheckValue = this.sSelectCheckValue.filter((itemIndex) => {
                return itemIndex != index;
            });
        } else {
            this.sSelectCheckValue.push(index);
        }
    }

    this.getShowSelectValueImage = function (id) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(id, this.sSelectCheckValue);
        }
        if (bValue) {
            return "/images/select_select.png";
        }
        return "/images/select_box.png";
    }

    this.onCheckSelectItem = function (id) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(id, this.sSelectCheckValue);
        }
        return bValue
    }

    this.onCloseCheck = function () {
        this.sSelectCheckValue = []
        this.bShowCheckDetail = false
    }

    this.onStartCheck = function (item) {
        if(this.sSelectCheckValue.length == 0){
            teaweb.warn("请选择体检项目")
            return
        }
        teaweb.confirm("确定立即体检吗？", function () {
            this.$post(".scans").params({
                Opt:'now',
                VirusPath:this.searchPath,
                WebShellPath:this.webSearchKey,
                MacCode: [this.MacCode],
                ScanItems: this.sSelectCheckValue.join(","),

            }).refresh()
        })

        this.bShowCheckDetail = false

        //
    }
    this.onStopCheck = function (item) {
        teaweb.confirm("确定取消体检吗？", function () {
            this.$post(".scans").params({
                Opt:'cancel',
                MacCode: [item.macCode],
            }).refresh()
        })
    }
    //检测是否显示扫描路径的输入框和提示框
    this.onCheckSelectValue = function () {

        var selextBox = document.getElementsByName("customScan")
        if (selextBox) {
            for (var item of selextBox) {
                if (item.checked) {
                    if (item.value == 2) {
                        this.bTimeOutTip = true
                        this.bShowScanPath = false
                    } else if (item.value == 3) {
                        this.bTimeOutTip = false
                        this.bShowScanPath = true
                    } else {
                        this.bTimeOutTip = false
                        this.bShowScanPath = false
                    }
                }
            }
        }
    }


    this.checkShowColor = function (curValue, maxValue) {
        if(curValue && maxValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            return tempValue >= 100
        }
        return false
    }

    this.getProgressPer = function (curValue, maxValue) {
        if(curValue && maxValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            return tempValue + "%"
        }
        return "0%"
    }

    //计时器
    this.testTime1 = null
    this.testTime2 = null
    //传入的值需要用 'testTime1' 命名
    this.onCreateTimeOut = function (timeId) {
        this.onReleaseTimeOut(timeId)
        this[timeId] = createTimer(function () {
        }, {timeout: 1000});
        this[timeId].start();
    }

    this.onReleaseTimeOut = function (timeId) {
        console.log(this[timeId])
        if (this[timeId]) {
            this[timeId].stop()
            this[timeId] = null
        }
    }

    this.tableData = [
        {
            id: 1,
            hostData: {
                intIp: "192.168.0.1",
                outIp: "192.168.1.1",
                hostName: "elk.novalocal",
                systemName: "centos linux7.6.1810_64bit",
                rootName: "3.10.0-957.1.3.el7",
                macAddr: "3.10.0-957.1.3.el7",
                remarks: "备注信息",
            },
            status: 3,
            checkNum: 70,
            startTime: "2021-06-05T12:15:25.000",
            endTime: "2021-06-05T13:15:25.000",
            maxValue: 100,
            curValue: 100
        },
        {
            id: 2,
            hostData: {
                intIp: "192.168.0.1",
                outIp: "192.168.1.1",
                hostName: "elk.novalocal",
                systemName: "centos linux7.6.1810_64bit",
                rootName: "3.10.0-957.1.3.el7",
                macAddr: "3.10.0-957.1.3.el7",
                remarks: "备注信息",
            },
            status: 2,
            checkNum: 70,
            startTime: "2021-06-05T12:15:25.000",
            endTime: "2021-06-05T13:15:25.000",
            maxValue: 100,
            curValue: 70
        }
    ]

})