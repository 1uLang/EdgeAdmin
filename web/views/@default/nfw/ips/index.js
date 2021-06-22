Tea.context(function () {
    this.severity=1



    this.getStatus= function (status) { 
        switch(status){
            case 1:
                return "已开启"
            case 2:
                return "已停用"  
            default:
                return "已停用"
        }
    }

    this.getEditName= function (bool) { 
        return bool?"drop":"alert";
    }

    this.getItemInfo = function (id) { 
        for (var i=0;i<this.tableData.length;i++){
            if(this.tableData[i].id ==id){
                return this.tableData[i]
            }
        }
        return null
    }

    //启用
    this.onOpenCig = function (id) {  }
    //禁用
    this.onCloseCig = function (id) {  }
    //更换alert/drop
    this.onChangeCig = function (id) { 
        var itemData = this.getItemInfo(id)
        teaweb.confirm("确定要删除所选规则？", function () {

        })
     }

    this.hostData = [
        {id:1,hostAddress:"成都-ddos-192.168.1.1",},
        {id:2,hostAddress:"成都-ddos-192.168.1.2",},
        {id:3,hostAddress:"成都-ddos-192.168.1.3",},
        {id:4,hostAddress:"成都-ddos-192.168.1.4",},
    ]


    this.tableData = [
        {id:1,value1:"900505001",value2:"drop",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:1,editType:1},
        {id:2,value1:"900505001",value2:"alert",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:2,editType:2},
        {id:3,value1:"900505001",value2:"drop",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:1,editType:1},
        {id:4,value1:"900505001",value2:"drop",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:2,editType:1},
        {id:5,value1:"900505001",value2:"drop",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:1,editType:1},
    ]

})