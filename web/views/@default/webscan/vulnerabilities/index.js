Tea.context(function () {
  this.onOpenDetail = function (itemId) {
    
    if (!itemId) {
      document.getElementById("rightDiv").style.display = "none";
    } else {
      document.getElementById("rightDiv").style.display = "flex";
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
});
