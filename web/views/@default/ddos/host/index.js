Tea.context(function () {
    this.severity = ''
    this.dropNode=''
    this.nShowState = 1 //当前显示的页面

    this.tableState=1 //只有在nShowState==2时生效 屏蔽和连接列表的切换

    this.onAddNodeIP = function () { 
        teaweb.popup(Tea.url(".createPopup"), {
            callback: function () {
              teaweb.success("保存成功", function () {
                teaweb.reload();
              });
            },
          });
    }
    //分页切换
    this.changeState = function(state){
      if(this.nShowState!=state){
        this.nShowState = state
      }
    }
    this.onOpenConfig = function (id) { 
      this.changeState(2)
    }

    //配置里面的列表切换
    this.changeListState = function(state){
      if(this.tableState!=state){
        this.tableState = state
      }
    }

    //查询
    this.onCheck = function () {
      var inputValue = ""
      if(this.tableState==1){
        inputValue = document.getElementById("linkHostInput").value
      }else{
        inputValue = document.getElementById("unlinkHostInput").value
      }
      console.log(inputValue)
      //todo
    }
    //导出
    this.onExport = function () {  
      var inputValue = ""
      if(this.tableState==1){
        inputValue = document.getElementById("linkHostInput").value
      }else{
        inputValue = document.getElementById("unlinkHostInput").value
      }
      console.log(inputValue)
      //todo
    }
    //全部释放 如果传入id 则单独释放 否则释放全部
    this.onRelease = function (id) { 

    }


    this.tableData = [
      {id:1,host:"118.112.240.0",user:"luobing",value1:"0.15",value2:"0.00",value3:"0.00",value4:"0.00",value5:"0.00",value6:"0.00",value7:"0.00",value8:"0.00",value9:"0.00",value10:"0.00",value11:"0.00",value12:"0.00",value13:"0.00"},
      {id:1,host:"118.112.240.0",user:"luobing",value1:"0.15",value2:"0.00",value3:"0.00",value4:"0.00",value5:"0.00",value6:"0.00",value7:"0.00",value8:"0.00",value9:"0.00",value10:"0.00",value11:"0.00",value12:"0.00",value13:"0.00"},
      {id:1,host:"118.112.240.0",user:"luobing",value1:"0.15",value2:"0.00",value3:"0.00",value4:"0.00",value5:"0.00",value6:"0.00",value7:"0.00",value8:"0.00",value9:"0.00",value10:"0.00",value11:"0.00",value12:"0.00",value13:"0.00"},
      {id:1,host:"118.112.240.0",user:"luobing",value1:"0.15",value2:"0.00",value3:"0.00",value4:"0.00",value5:"0.00",value6:"0.00",value7:"0.00",value8:"0.00",value9:"0.00",value10:"0.00",value11:"0.00",value12:"0.00",value13:"0.00"},
    ]

    this.linkData=[
      {id:1,host:"118.112.240.0",user:"118.112.240.0",status:"禁止689883471秒",res:"集群同步屏蔽 + 系统连接保护"},
      {id:1,host:"118.112.240.0",user:"118.112.240.0",status:"禁止689883471秒",res:"集群同步屏蔽 + 系统连接保护"},
      {id:1,host:"118.112.240.0",user:"118.112.240.0",status:"禁止689883471秒",res:"集群同步屏蔽 + 系统连接保护"},
    ]

})