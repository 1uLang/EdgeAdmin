Tea.context(function () {

    this.keyword = ""
    this.curIndex = -1
    this.onGoBack = function(){
        window.location = "/hids/invade"
    }
    this.onOpenDetail = function(id){
        
    }

    this.onGoBack = function(){
        window.location = "/hids/invade/virus"
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
          var tempTime = time.substring(0, time.indexOf("."));
          resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
      };

    this.pageData = {
        id:1,
        systemInfo: "centos linux7.6.1810_64bit",
        lastScanTime: "2021-06-07T11:33:06.000",
        countName:"病毒木马数",
        countValue:2,
        tableTitle:[
            {name:"主机IP",width:"834px"},
            {name:"病毒木马数",width:"200px"},
            {name:"详情",width:"90px"}
        ],
        menuData:[
            {id:1,value:"待处理（12）"},
            {id:2,value:"已处理（8）"},
        ],
        tableTitle:[
            {id:1,titles:["异常登录类型","登录信息","发现时间","来源","状态","操作"]}
        ],
        //请求相关的配置
        reqMethon:"get",    
        reqPath:".../...",  

    }
})