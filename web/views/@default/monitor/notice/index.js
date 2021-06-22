Tea.context(function () {
    this.keyword = ""


    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
          var tempTime = time.substring(0, time.indexOf("."));
          resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
      };

    this.onDelete = function (id) {  

    }

    this.tableData = [
        {id:1,warringDesc:"您的URL：https://www.fapiao.com 状态异常，请及时处理！",user:"luobing",warringTime:"2021-02-06T15:12:26.000",},
        {id:2,warringDesc:"您的URL：https://www.fapiao.com 状态异常，请及时处理！",user:"lusir",warringTime:"2021-03-06T15:12:26.000",},
        {id:3,warringDesc:"您的URL：https://www.fapiao.com 状态异常，请及时处理！",user:"dengp",warringTime:"2021-05-06T15:12:26.000",}
    ]
})