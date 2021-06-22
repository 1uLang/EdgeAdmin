Tea.context(function () {
    this.severity=1

    this.nShowState = 1

    this.editId = 1             //操作的ID
    this.postId = 1             //接口* ID
    this.dirctionId = 1         //方向 *ID
    this.agreementTypeId = 1    //TCP/IP版本*
    this.agreementId = 1        //协议*
    this.sourceTypeId = 1       //源 单机网络的id
    this.sourceValue = "any"    //源的值
    this.sourceId = 1           //源的ID
    this.editSourceId = 1       //源修改界面需要下拉选择
    this.targetTypeId = 1       //目标 单机网络的id
    this.targetValue = "any"    //目标的值
    this.targetId = 1           //目标的ID
    this.editTargetId = 1       //目标修改界面需要下拉选择
    this.descValue = ""         //描述值
    this.targerRangeStartId = 1 //修改中 目标范围的开始值
    this.targerRangeEndId = 1 //修改中 目标范围的终止值

    this.onChangeShowState=function (state) { 
        if(this.nShowState != state){
            this.nShowState = state
            if(state ==2){
                this.resetValue()
            }
        }
    }

    this.getStatus= function (status) { 
        return status == 1?"已开启":"已停用";
    }

    this.getDirection= function (type) { 
        return type == 1?"进":"出";
    }


    this.getTactics= function (type) { 
        switch(status){
            case 1:
                return "通过"
            case 2:
                return "阻止"  
            case 3:
                return "拒绝"
            default:
                return "拒绝"
        }
    }

    this.getItemInfo = function (id) { 
        for (var i=0;i<this.tableData.length;i++){
            if(this.tableData[i].id ==id){
                return this.tableData[i]
            }
        }
        return null
    }

    //修改配置
    this.onOpenChangeView=function (id) { 
        console.log(id)
        for (var i=0;i<this.tableData.length;i++)
        { 
            if(this.tableData[i].id == id){
                this.editId = 1             
                this.postId = 1             
                this.dirctionId = 1         
                this.agreementTypeId = 1    
                this.agreementId = 1        
                // this.sourceTypeId = 1    //不让修改   
                this.sourceValue = "赋值"    
                this.sourceId = 1       
                this.editSourceId = 1      
                // this.targetTypeId = 1    //不让修改  
                this.targetValue = "赋值"        
                this.targetId = 1           
                this.descValue = "描述" 
                this.editTargetId = 1   
                this.targerRangeStartId = 1
                this.targerRangeEndId = 1     

                break
            }
        }
        this.onChangeShowState(3)
    }

    //重置value
    this.resetValue = function () { 
        this.editId = 1             
        this.postId = 1             
        this.dirctionId = 1         
        this.agreementTypeId = 1    
        this.agreementId = 1        
        this.sourceTypeId = 1       
        this.sourceValue = "any"    
        this.sourceId = 1           
        this.targetTypeId = 1       
        this.targetValue = "any"      
        this.targetId = 1     
        this.editSourceId = 1    
        this.editTargetId = 1         
        this.descValue = ""     
        this.targerRangeStartId = 1
        this.targerRangeEndId = 1   
     }

    //保存配置
    this.onSaveConfig= function () { 

    }

    //开启配置
    this.onOpenConfig = function (id) { 

    }

    //关闭配置
    this.onCloseConfig = function (id) {

    }
    //删除配置
    this.onDelConfig = function (id) {
        teaweb.confirm("确定要删除所选规则？", function () {

        })
    }


    //源和目标的位数
    this.dropNodeData = [
        {id:1,value:1},
        {id:2,value:2},
        {id:3,value:3},
        {id:4,value:4},
        {id:5,value:5},
        {id:6,value:6},
        {id:7,value:7},
        {id:8,value:8},
    ]

    //操作
    this.editData = [
        {id:1,value:"通过"},
        {id:2,value:"拒绝"},
        {id:3,value:"阻止"}
    ]
    this.dirctionData = [
        {id:1,value:"in"},
        {id:2,value:"out"},
    ]

    this.hostData = [
        {id:1,hostAddress:"成都-ddos-192.168.1.1",},
        {id:2,hostAddress:"成都-ddos-192.168.1.2",},
        {id:3,hostAddress:"成都-ddos-192.168.1.3",},
        {id:4,hostAddress:"成都-ddos-192.168.1.4",},
    ]
    

    //postIndex:接口* 选择的id  typeIndex 类型* 选择的id  sourceTypeIndex 源* 选择的id  targetIndex 目标* 选择的id
    this.tableData = [
        {id:1,postId:"LAN",directionType:1,agreement:"IPv4 TCP",value1:"192.168.0.1",post1:"8877",value2:"192.168.1.1",post2:"3306",desc:"描述",tactics:1,status:1,postIndex:1,typeIndex:1,sourceTypeIndex:1,targetIndex:1,sourceValue:"源的值"},
        {id:2,postId:"WAN",directionType:2,agreement:"IPv4 TCP",value1:"192.168.0.1",post1:"8877",value2:"192.168.1.1",post2:"3306",desc:"描述",tactics:1,status:1,postIndex:1,typeIndex:1,sourceTypeIndex:1,targetIndex:1,sourceValue:"源的值"},
        {id:3,postId:"Loopback",directionType:1,agreement:"IPv4 TCP",value1:"192.168.0.1",post1:"8877",value2:"192.168.1.1",post2:"3306",desc:"描述",tactics:1,status:1,postIndex:1,typeIndex:1,sourceTypeIndex:1,targetIndex:1,sourceValue:"源的值"},
    ]

})