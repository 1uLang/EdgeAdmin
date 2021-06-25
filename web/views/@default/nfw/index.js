Tea.context(function () {
    this.severity=1

    this.strCPUPer = "0%"
    this.strStatusPer = "0%"
    this.strMBUFPer = "0%"
    this.strRAMPer = "0%"
    this.strSWAPPer = "0%"
    this.strDiskPer = "0%"

    this.$delay(function () {
        this.initDataPer()
    })

    this.initDataPer = function(){
        this.strCPUPer = this.tableData.cpuPer + "%"
        this.strStatusPer = this.getProgressPer(this.tableData.statusCount.maxValue,this.tableData.statusCount.useValue)
        this.strMBUFPer = this.getProgressPer(this.tableData.MBUFCount.maxValue,this.tableData.MBUFCount.useValue)
        this.strRAMPer = this.getProgressPer(this.tableData.RAMPer.maxValue,this.tableData.RAMPer.useValue)
        this.strSWAPPer = this.getProgressPer(this.tableData.SWAPPer.maxValue,this.tableData.SWAPPer.useValue)
        this.strDiskPer = this.tableData.diskPer + "%"

        document.getElementById("CPUItem").style.width = this.strCPUPer;
        document.getElementById("statusItem").style.width = this.strStatusPer;
        document.getElementById("MBUFItem").style.width = this.strMBUFPer;
        document.getElementById("RAMItem").style.width = this.strRAMPer;
        document.getElementById("SWAPItem").style.width = this.strSWAPPer;
        document.getElementById("diskItem").style.width = this.strDiskPer;
        
    }

    //时间转换 如果需要用的话
    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
          var tempTime = time.substring(0, time.indexOf("."));
          resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
      };

    this.getProgressPer =function (maxValue,curValue) {
        return ((curValue/maxValue)*100).toFixed(1) + "%"
    }

     this.getShowPerText = function(strPer,maxValue,curValue,unit){
        let showValue = strPer + "("+curValue+"/"+maxValue
        if(unit){
            showValue = showValue+unit
        }
        showValue = showValue + ")"
        return showValue
     }
        

    this.hostData = [
        {id:1,hostAddress:"成都-ddos-192.168.1.1",},
        {id:2,hostAddress:"成都-ddos-192.168.1.2",},
        {id:3,hostAddress:"成都-ddos-192.168.1.3",},
        {id:4,hostAddress:"成都-ddos-192.168.1.4",},
    ]

    this.tableDatas = {
        verInfo:"CloudFW 21.1.5-amd64<br />FreeBSD 12.1-RELEASE-p16-HBSD<br />OpenSSL 1.1.1k 25 Mar 2021",
        cpuInfo:"CPU type",
        slbInfo:"0.41, 0.34, 0.27",
        runTime:"12:30:00",
        timeNow:"Sun Jun 6 16:25:34 UTC 2021",
        lastConfigTime:"Sun Jun 6 16:23:36 UTC 2021",
        cpuPer:30,
        statusCount:{maxValue:8000,useValue:1000},
        MBUFCount:{maxValue:13000,useValue:500},
        RAMPer:{maxValue:2048,useValue:1024,unit:"MB"},
        SWAPPer:{maxValue:2048,useValue:512,unit:"MB"},
        diskPer:30,
    }
})