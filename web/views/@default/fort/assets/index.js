Tea.context(function () {
    
    this.curIndex = -1

    this.pageState = 1

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
    this.getStatus = function (status) {
        switch (status) {
            case 1:
                return "启已用"
            case 0:
                return "已停用"
            default:
                return "已停用"
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
    this.onEnabled = function (id) { 

    }
    this.onDisabled =function (id) {

    }
    this.onDelete = function (id) { 
        
     }
 
    this.tableData = [
        {id:1,value1:"智安安全审计系统服务器",value2:"47.108.234.195",value3:"8 Core 7.82 G 49.0 G",value4:1,value5:1,value6:"2021-06-16T09:29:23.140"},
        {id:2,value1:"智安安全审计系统服务器",value2:"47.108.234.195",value3:"8 Core 7.82 G 49.0 G",value4:2,value5:0,value6:"2021-06-16T09:29:23.140"},
        {id:3,value1:"智安安全审计系统服务器",value2:"47.108.234.195",value3:"8 Core 7.82 G 49.0 G",value4:0,value5:1,value6:"2021-06-16T09:29:23.140"},
        {id:4,value1:"智安安全审计系统服务器",value2:"47.108.234.195",value3:"8 Core 7.82 G 49.0 G",value4:1,value5:0,value6:"2021-06-16T09:29:23.140"},
    ]
})