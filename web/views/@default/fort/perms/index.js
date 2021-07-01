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

    this.tableData= [
        {id:1,value1:"安全审计系统服务器授权",value2:"用户",value3:"1",value4:"1",value5:1},
        {id:2,value1:"安全审计系统服务器授权",value2:"用户",value3:"1",value4:"1",value5:0},
        {id:3,value1:"安全审计系统服务器授权",value2:"用户",value3:"1",value4:"1",value5:1},
    ]

    this.baseInfo = {   id:1,value:"42f167c2-d91a-4f20-99b1-3d56dabd896a",value2:"智安审计服务器授权",value3:3,value4:1,value5:1,
            value6:"全部,连接,上传文件,下载文件,上传下载,剪切板复制,剪切板粘贴,复制粘贴",value7:"2021-06-05T18:20:50.000",
            value8:"2021-06-05T18:20:50.000",value9:"2021-06-05T18:20:50.000",value10:"Administrator",value11:"这是备注信息"
        }
    
})