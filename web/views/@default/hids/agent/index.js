Tea.context(function () {


    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
          var tempTime = time.substring(0, time.indexOf("."));
          resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
      };


    this.onOpenCommand = function(){
        teaweb.popup(Tea.url(".createCommand"));
    }

    this.onStartConfig = function(id){
        
    }

    this.onStopConfig = function(id){

    }

    this.onDelete = function(id){
        teaweb.confirm("确定要删除？", function () {
           
        })
    }

    this.getStatusImgName = function(status){
        return status == 1? "已启用":"已停用"
    }

    this.tableData = [
        {id:1,ip:"192.168.0.1",status:1,edition:"(Windows)5.4.30014",installTime:"2021-06-07T11:33:06.000",latelyOnlineTime:"2021-06-08T11:33:06.000",lastOnlineTime:"2021-06-09T11:33:06.000"},
        {id:2,ip:"192.168.0.2",status:2,edition:"(Linux64)4.0.30693",installTime:"2021-06-07T11:33:06.000",latelyOnlineTime:"2021-06-08T11:33:06.000",lastOnlineTime:"2021-06-09T11:33:06.000"}
    ]
});
  