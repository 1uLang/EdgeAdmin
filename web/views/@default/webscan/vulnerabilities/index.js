Tea.context(function () {
  this.address = ''
  this.severity = ''
  this.detailInfo = null
  this.onOpenDetail = function (vul) {
      if(vul){
        detailInfo = this.$get(".details")
        .params({
          VulId: vul.vuln_id
        })
        document.getElementById("rightDiv").style.display = "flex";
      }else{
        document.getElementById("rightDiv").style.display = "none";
      }
  };

  this.onChangeTimeFormat = function (time) {
    var resultTime = "";
    if (time) {
      var tempTime = time.substring(0, time.indexOf("."));
      resultTime = tempTime.replace("T", " ");
    }
    return resultTime;
  };
  this.onChangeSeverityFormat = function (severity) {
    var resultSeverity = severity;
    if (severity) {
      switch (severity) {
        case 3:return  '高危'
        case 2:return  '中危'
        case 1:return  '低危'
        default:return  '信息'
      }
    }
    return resultSeverity;
  };
});
