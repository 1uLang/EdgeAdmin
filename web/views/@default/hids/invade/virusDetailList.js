Tea.context(function () {

    this.keyword = ""
    this.nPageType = -1
    this.sShowMenuData = []     //分页数组  [{name:"",state:1}]
    this.nShowTableState = -1   //当前选中的分页的state
    this.nShowValueCount = 0    //table值的
    this.pShowTableTile = null  //显示的table 表头
    this.pShowTableData = null  //显示的table的value数组
    this.sTableValueKey =[]     //table value的key值 [string]
    this.sShowBtnData = []      //按钮数组 [{name:"",funcName:""}]

    this.bShowItemDetail = false    //详情的显隐
    this.pClickItemData = null  //结构 { width:,height:,itemData:[{name:"",value:""}]}


    this.sMenuData = [
        {id:1,value:"待处理（count）"},
        {id:2,value:"已处理（count）"},
        {id:3,value:"已处理（count）"}
    ]

    this.sBtnData = [
        {id:1,value:"详情"},
        {id:2,value:"关闭"},
        {id:3,value:"取消关闭"},
    ]

    this.$delay(function () {
        this.onOpenPageDetail(1)
    })


    this.onChangeTableStat = function(state){
        if(this.nShowTableState!=state){
            this.nShowTableState = state
            this.onOpenPageDetail(state)
        }
    }

    this.onGoBack = function(){
        window.location = "/hids/invade"
    }

    //打开页面 刷新页面数据
    this.onOpenPageDetail = function(id){
        this.nShowTableState = id
        this.onGetTableInfo(id)

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

      //获取分页按钮名称
      this.getMenuName = function(id,itemCount){
          if(!itemCount){
            itemCount = 0
          }
          for(var i=0;i<this.sMenuData.length;i++){
              if(this.sMenuData[i].id == id){
                  let menuName = this.sMenuData[i].value.replace("count", itemCount)
                  let menuData = {state:id,name:menuName}
                return menuData;
              }
          }
          return null
      }
      
      //匹配表格内部按钮名称
      this.getTableBtnName = function(id){
        for(var i=0;i<this.sBtnData.length;i++){
            if(this.sBtnData[i].id == id){
              return this.sBtnData[i].value;
            }
        }
        return null
      }
      //获取表格内部按钮数据
      this.onGetTableBtnData = function(btnData){
          if(btnData){
            for(var i = 0;i<btnData.length;i++){
                let btnName = this.getTableBtnName(btnData[i].id)
                if(btnName){
                    let curBtnData = {name:btnName,funcName:btnData[i].func}
                    this.sShowBtnData.push(curBtnData)
                }
                
            }
          }
      }

      //获取table数据
      this.onGetTableInfo = function (id) { 
        this.sShowMenuData = []
        this.sTableValueKey =[]
        this.sShowBtnData = [] 
        this.pShowTableTile = null
        this.pShowTableData = null
        this.nShowValueCount = 0
        this.nPageType = this.pageData.pageType
        if(this.pageData.tableData){
            //获取表格数据
          for (var i = 0; i < this.pageData.tableData.length; i++) {
            if (this.pageData.tableData[i].id == id) {
                this.pShowTableTile = this.pageData.tableData[i].titleData
                this.nShowValueCount = this.pageData.tableData[i].titleData.length-1 //最后一个是操作 所以需要减1
                this.pShowTableData = this.pageData.tableData[i].titleValue
                //获取btn数据
                this.onGetTableBtnData(this.pageData.tableData[i].btnData)
                break
            }
          }
          //获取tableValue的字段名称
          if(this.pShowTableTile){
              for(var item of this.pShowTableTile){
                  if(item.key){
                    this.sTableValueKey.push(item.key)
                  }
              }
          }
          //获取分页数据
          for(var i = 0;i < this.pageData.tableData.length; i++){
            let value = this.getMenuName(this.pageData.tableData[i].menuId,this.pageData.tableData[i].titleValue.length)
            this.sShowMenuData.push(value)
          }
        }
      }

    this.getStatusName = function(status){
        switch(status){
            case 1:
                return "待处理"
            case 2:
                return "处理中"
            case 3 :
                return "已处理"
            case 4:
                return "异常"
            default :
                return "未知"
        }
    }

    this.getTableValue = function(key,value){
        if(key.indexOf("time") != -1){
            return this.onChangeTimeFormat(value)
        }else if(key.indexOf("status") != -1){
            return this.getStatusName(value)
        }else{
            return value
        }
    }

    //执行方法
    this.onRunFunction = function(funcName,itemId){
        console.log(itemId)
        switch(funcName){
            case "openDetail":
                this.onOpenDetail(itemId)
                break
            case "close":
                this.onClose(itemId)
                break;
            case "open":
                this.onOpen(itemId)
                break 
            default :
                break
        }
    }

    this.onCloseDetail = function(){
        this.bShowItemDetail = false
    }
    //打开详情
    this.onOpenDetail = function(id){
        console.log("onOpenDetail")
        this.bShowItemDetail = true
        //通过req获取详情
        this.pClickItemData = this.itemDetailInfo

        // switch(this.nPageType){ //根据页面类型来执行不同的流程

        // }
    }
    //打开
    this.onOpen = function(id){
        console.log("onOpen")
        // switch(this.nPageType){ //根据页面类型来执行不同的流程

        // }
    }
    //关闭
    this.onClose = function(id){
        console.log("onClose")
        // switch(this.nPageType){ //根据页面类型来执行不同的流程

        // }
    }

    this.itemDetailInfo = {
        width:"800px",
        height:"500px",
        itemInfo:[
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"},
            {key:"主机IP",value:"156.240.95.243(172.31.13.84)"}
        ]
    }

    this.pageData = {
        id:1,
        pageType:1,     //页面类型
        //标题模版
        systemInfo: "centos linux7.6.1810_64bit",
        lastScanTime: "2021-06-07T11:33:06.000",
        countName:"病毒木马数",
        countValue:2,
        //表格数据
        //表头
        tableData:[
            {   id:1,
                menuId:1,
                btnData:[{id:1,func:"openDetail"},{id:2,func:"close"}],
                titleData:[
                    {name:"异常登录类型",width:"174px",key:"exceptionType"},
                    {name:"登录信息",width:"371px",key:"loginInfo"},
                    {name:"发现时间",width:"190px",key:"time"},
                    {name:"来源",width:"90px",key:"comeSource"},
                    {name:"状态",width:"124px",key:"status"},
                    {name:"操作",width:"191px"}
                ],
                titleValue:[
                    {id:1,exceptionType:"异常 IP 段登录",loginInfo:"root 登录 IP：192.168.145.1",time:"2021-06-07T16:28:04.000",comeSource:"本地扫描",status:1}
                ],
            },
            {   id:2,
                menuId:2,
                btnData:[{id:1,func:"openDetail"},{id:3,func:"open"}],
                titleData:[
                    {name:"异常登录类型",width:"174px",key:"exceptionType"},
                    {name:"登录信息",width:"371px",key:"loginInfo"},
                    {name:"发现时间",width:"190px",key:"time"},
                    {name:"来源",width:"90px",key:"comeSource"},
                    {name:"状态",width:"124px",key:"status"},
                    {name:"操作",width:"191px"}
                ],
                titleValue:[
                    {id:1,exceptionType:"异常 IP 段登录123",loginInfo:"root 登录 IP：192.168.145.1  123",time:"2021-06-07T16:28:04.000",comeSource:"本地扫描123",status:2}
                ]
            },
            
        ],

        //请求相关的配置
        reqMethon:"get",    
        reqPath:".../...",  

    }
})