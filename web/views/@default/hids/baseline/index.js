Tea.context(function () {

    this.nCheckState = 1    //检测状态
    this.nResultState = 1   //检查结论
    this.dayFrom = ""
    this.dayTo = ""
    this.curIndex = -1

    this.bShowCheckDetail = false
    this.pCheckDetailData = null
    this.sSelectValue = []

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })

    this.onChangeCheckState = function (index) { 
        if(this.nCheckState!=index){
            this.nCheckState = index
        }
     }

    this.onChangeResultState = function(index){
        if(this.nResultState!=index){
            this.nResultState = index
        }
    }

    this.onChangeTimeState = function () { 

     }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "-";
        if (time && time.length > 0) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.getStatusName = function (status) { 
        switch(status){
            case 1:
                return "未检查"
            case 2:
                return "检查中"
            case 3:
                return "已完成"
            default:
                return "未知"
        }
     }

     this.enters = function (index) {
        this.curIndex = index;
      }
    
      this.leaver = function (index) {
        this.curIndex = -1;
      }
    
    this.getProgressPer = function (curValue,maxValue) { 
        if(curValue && maxValue && maxValue >= curValue){
            return ((curValue/maxValue)*100).toFixed(1)+"%"
        }
        return "0%"
     }

     this.getItemInfo = function (id) { 
        this.pCheckDetailData = null
        for (var i = 0; i < this.tableData.length; i++) {
            if (this.tableData[i].id == id) {
                this.pCheckDetailData = this.tableData[i]
                break
            }
        }
      }

     this.onOpenCheck = function (id) { 
         this.getItemInfo(id)
         if(this.pCheckDetailData){
            this.bShowCheckDetail =  true
         }else{
             this.bShowCheckDetail = false
         }
        
      }

      this.onCloseCheck = function () { 
        this.bShowCheckDetail = false
       }

    this.onStartCheck = function (id) {
        this.onCloseCheck()
    }

    this.onStopCheck = function (id) { 

    }

    this.onOpenDetail = function (id) { 
        window.location = "/hids/baseline/detail"
    }

    //检测是否包含元素
    this.checkSelectValue = function (index) {
        for (var i = 0; i < this.sSelectValue.length; i++) {
            if (this.sSelectValue[i] == index) {
                return true;
            }
        }
        return false;
    }

    //添加/删除元素
    this.onAddSelectValue = function (index) {
        let bValue = false;
        if(this.checkSelectValue){
            bValue = this.checkSelectValue(index);
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
        if(this.sSelectValue){
          bValue = this.checkSelectValue(id);
        }
        if (bValue) {
          return "/images/select_select.png";
        }
        return "/images/select_box.png";
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
            checkValue:[
                {id:1,value:"Linux等保三级 (LINUX)"},
                {id:2,value:"Linux等保三级 (LINUX)"},
                {id:3,value:"Linux等保三级 (LINUX)"},
                {id:4,value:"Linux等保三级 (LINUX)"},
                {id:5,value:"Linux等保三级 (LINUX)"},
                {id:6,value:"Linux等保三级 (LINUX)"},
                {id:7,value:"Linux等保三级 (LINUX)"},
                {id:8,value:"Linux等保三级 (LINUX)"},
                {id:9,value:"Linux等保三级 (LINUX)"},
            ],
            status: 1,
            finishTime: "",
            maxCount:100,
            curCount:10,
            checkAllCount:0,
            checkCurCount:0,
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
            checkValue:[
                {id:1,value:"Linux等保三级 (LINUX)"},
                {id:2,value:"Linux等保三级 (LINUX)"},
                {id:3,value:"Linux等保三级 (LINUX)"},
                {id:4,value:"Linux等保三级 (LINUX)"},
                {id:5,value:"Linux等保三级 (LINUX)"},
                {id:6,value:"Linux等保三级 (LINUX)"},
                {id:7,value:"Linux等保三级 (LINUX)"},
                {id:8,value:"Linux等保三级 (LINUX)"},
                {id:9,value:"Linux等保三级 (LINUX)"},
            ],
            status: 2,
            finishTime: "",
            maxCount:100,
            curCount:30,
            checkAllCount:11,
            checkCurCount:1,
          },
          {
            id: 3,
            hostData: {
              intIp: "192.168.0.1",
              outIp: "192.168.1.1",
              hostName: "elk.novalocal",
              systemName: "centos linux7.6.1810_64bit",
              rootName: "3.10.0-957.1.3.el7",
              macAddr: "3.10.0-957.1.3.el7",
              remarks: "备注信息",
            },
            checkValue:[
                {id:1,value:"Linux等保三级 (LINUX)"},
                {id:2,value:"Linux等保三级 (LINUX)"},
                {id:3,value:"Linux等保三级 (LINUX)"},
                {id:4,value:"Linux等保三级 (LINUX)"},
                {id:5,value:"Linux等保三级 (LINUX)"},
                {id:6,value:"Linux等保三级 (LINUX)"},
                {id:7,value:"Linux等保三级 (LINUX)"},
                {id:8,value:"Linux等保三级 (LINUX)"},
                {id:9,value:"Linux等保三级 (LINUX)"},
            ],
            status: 3,
            finishTime: "2021-06-07T11:33:06.000",
            maxCount:100,
            curCount:70,
            checkAllCount:11,
            checkCurCount:0,
          },
    ]

});
  