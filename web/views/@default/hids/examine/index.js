Tea.context(function () {

    this.nCheckState = 1    //体检状态
    this.nResultState = 1   //体检分数
    this.nHealthNumState = 1      //分数类型
    this.sSelectValue = []  //体检项目
    this.dayFrom = ""
    this.dayTo = ""

    this.curIndex = -1

    this.bShowCheckDetail = false
    this.pCheckDetailData = null
    this.sSelectCheckValue = []

    this.webSearchKey = ""
    this.searchPath = ""

    this.bTimeOutTip = false
    this.bShowScanPath = false

    this.sTopSelectItem = [
        {id:1,value:"系统漏洞"},
        {id:2,value:"弱口令"},
        {id:3,value:"风险账号"},
        {id:4,value:"配置缺陷"},
        {id:5,value:"病毒木马"},
        {id:6,value:"网页后门"}
    ]
    this.sBottomSelectItem = [
        {id:7,value:"反弹shell"},
        {id:8,value:"异常账号"},
        {id:9,value:"系统命令篡改"},
        {id:10,value:"异常进程"},
        {id:11,value:"日志异常删除"},
    ]

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })

    this.onChangeCheckState = function(state){
        if(this.nCheckState!=state){
            this.nCheckState = state
        }
    }

    this.onChangeHealthNumState = function(state){
        if(this.nHealthNumState!=state){
            this.nHealthNumState = state
        }
    }

    this.onChangeResultState = function(state){
        if(this.nResultState!=state){
            this.nResultState = state
        }
    }

    //检测是否包含元素
    this.checkSelectValue = function (index,selectValue) {
        if(selectValue){
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
        if(this.checkSelectValue){
            bValue = this.checkSelectValue(index,this.sSelectValue);
        }
        if (bValue) {
            this.sSelectValue = this.sSelectValue.filter((itemIndex) => {
                return itemIndex != index;
            });
        } else {
            this.sSelectValue.push(index);
        }
    }


    this.getShowSelectImage = function (id) {
        let bValue = false;
        if(this.checkSelectValue){
            bValue = this.checkSelectValue(id,this.sSelectValue);
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
        switch(status){
            case 1:
                return "未体检"
            case 2:
                return "体检中"
            case 3:
                return "已完成"
            default:
                return "未知"
        }
    }

    this.getCheckNumName = function (num) { 
        return (num && num>0) ? num+"分" : "未得出"
    }

    this.enters = function (index) {
        this.curIndex = index;
      }
    
      this.leaver = function (index) {
        this.curIndex = -1;
      }

    this.onOpenDetail = function (id) {
        window.location="/hids/examine/detail"
    }

    this.onOpenCheck = function (id) {
        this.sSelectCheckValue = []
        // .success(
            this.pCheckDetailData = [
                {
                    checkName:"漏洞风险检查项：",
                    checkValue:[
                        {id:1,value:"系统漏洞"},
                        {id:2,value:"系统漏洞"},
                        {id:3,value:"系统漏洞"},
                        {id:4,value:"系统漏洞"},
                        {id:5,value:"系统漏洞"},
                    ]
                },
                {
                    checkName:"入侵威胁检查项：",
                    checkValue:[
                        {id:6,value:"反弹shell"},
                        {id:7,value:"异常账号"},
                        {id:8,value:"系统命令篡改"},
                        {id:9,value:"异常进程"},
                        {id:10,value:"日志异常删除"},
                    ]
                }
            ]
            if(this.pCheckDetailData){
                this.bShowCheckDetail =  true
             }else{
                 this.bShowCheckDetail = false
             }
        // )
    }

    this.onSelectCheckValue = function(index) {
        let bValue = false;
        if(this.checkSelectValue){
            bValue = this.checkSelectValue(index,this.sSelectCheckValue);
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
        if(this.checkSelectValue){
            bValue = this.checkSelectValue(id,this.sSelectCheckValue);
        }
        if (bValue) {
            return "/images/select_select.png";
        }
        return "/images/select_box.png";
    }

    this.onCheckSelectItem = function (id) {
        let bValue = false;
        if(this.checkSelectValue){
            bValue = this.checkSelectValue(id,this.sSelectCheckValue);
        }
        return bValue
    }

    this.onCloseCheck = function () { 
        this.sSelectCheckValue = []
        this.bShowCheckDetail = false
    }

    this.onStartCheck = function (id) { 
        this.bShowCheckDetail = false
        //
    }
    this.onStopCheck = function (id) { 
        teaweb.confirm("确定取消体检吗？", function () {
			
		})
    }  
     //检测是否显示扫描路径的输入框和提示框
    this.onCheckSelectValue=function() {
        
        var selextBox = document.getElementsByName("customScan")
        if(selextBox){
            for(var item of selextBox){
                if(item.checked){
                    if(item.value==2){
                        this.bTimeOutTip = true
                        this.bShowScanPath = false
                    }else if(item.value==3){
                        this.bTimeOutTip = false
                        this.bShowScanPath = true
                    }else{
                        this.bTimeOutTip = false
                        this.bShowScanPath = false
                    }
                }
            }
        }
    }


     this.checkShowColor =function (curValue,maxValue) {
        var curValue = ((curValue/maxValue)*100).toFixed(1)
        console.log(curValue>=100);
        return curValue>=100
     }

    this.getProgressPer = function (curValue,maxValue) { 
        var curValue = ((curValue/maxValue)*100).toFixed(1)
        return curValue+"%"
    }

    this.tableData = [
        {   
            id:1,
            hostData: {
                intIp: "192.168.0.1",
                outIp: "192.168.1.1",
                hostName: "elk.novalocal",
                systemName: "centos linux7.6.1810_64bit",
                rootName: "3.10.0-957.1.3.el7",
                macAddr: "3.10.0-957.1.3.el7",
                remarks: "备注信息",
            },
            status:3,
            checkNum:70,
            startTime:"2021-06-05T12:15:25.000",
            endTime:"2021-06-05T13:15:25.000",
            maxValue:100,
            curValue:100
        },
        {   
            id:2,
            hostData: {
                intIp: "192.168.0.1",
                outIp: "192.168.1.1",
                hostName: "elk.novalocal",
                systemName: "centos linux7.6.1810_64bit",
                rootName: "3.10.0-957.1.3.el7",
                macAddr: "3.10.0-957.1.3.el7",
                remarks: "备注信息",
            },
            status:2,
            checkNum:70,
            startTime:"2021-06-05T12:15:25.000",
            endTime:"2021-06-05T13:15:25.000",
            maxValue:100,
            curValue:70
        }
    ]

})