Tea.context(function () {
    
    this.curIndex = -1

    this.pageState = 1

    this.mouseLeave = function () { 
        this.curIndex = -1
    }

    this.mouseEnter = function (index) { 
        this.curIndex = index
    }
    this.onChangeState = function (id) { 
        if( this.pageState!=id) {
            this.pageState = id
        }
    }

    this.onOpenDetail = function (id) { 
        this.onChangeState(3)
     }

    this.onEdit = function (id) { 

    }
    this.onDelete = function (id) {
        
    }

    //资产列表
    //推送
    this.onPush = function (id) { 

    }
    //删除
    this.onDeleteSource = function (id) {
        
    }

    this.getLinkStatus = function (status) { 
        switch (status) {
            case 1:
                return "可连接"
            case 0:
                return "不可连接"
            default:
                return "未知"
        }
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
          var tempTime = time.substring(0, time.indexOf("."));
          resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
      };

    this.tableData1=[
        {id:1,value1:"等保云demo服务器root账号",value2:"dby_web",value3:"ssh",value4:"自动登录",value5:"2",value6:1},
        {id:2,value1:"等保云demo服务器root账号",value2:"dby_web",value3:"ssh",value4:"自动登录",value5:"2",value6:2},
        {id:3,value1:"等保云demo服务器root账号",value2:"dby_web",value3:"ssh",value4:"自动登录",value5:"2",value6:0}
    ]

    this.tableData2=[
        {id:1,value1:"智安-安全审计系统服务器",value2:"182.150.0.104",value3:"等保云web(dby_web)"},
        {id:2,value1:"智安-安全审计系统服务器",value2:"182.150.0.104",value3:"等保云web(dby_web)"},
        {id:3,value1:"智安-安全审计系统服务器",value2:"182.150.0.104",value3:"等保云web(dby_web)"}
    ]

    this.baseInfo={
        id:1,value1:"42f167c2-d91a-4f20-99b1-3d56dabd896a",
        value2:"web",value3:"user",value4:"自动登录",value5:"/bin/whoaml",
        value6:"/bin/bash",value7:"2021-06-24T12:20:50.361",value8:"Administrator"
    }
})