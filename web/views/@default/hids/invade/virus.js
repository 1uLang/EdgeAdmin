Tea.context(function () {

    this.keyword = ""
    this.curIndex = -1
    this.onGoBack = function(){
        window.location = "/hids/invade"
    }
    this.onOpenDetail = function(id){
        window.location = "/hids/invade/virusDetailList"
    }
    this.enters = function (index) {
        this.curIndex = index;
      }
    
      this.leaver = function () {
        this.curIndex = -1;
      }

    this.pageData = {
        id:1,
        pageName:"病毒木马",
        tableTitle:[
            {name:"主机IP",width:"834px"},
            {name:"病毒木马数",width:"200px"},
            {name:"详情",width:"90px"}
        ],
        hostData:[
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
                postId:"192.168.0.1",
                count:2
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
                postId:"192.168.0.1",
                count:2
              },
        ],
        //请求相关的配置
        reqMethon:"get",    
        reqPath:".../...",  

    }
})