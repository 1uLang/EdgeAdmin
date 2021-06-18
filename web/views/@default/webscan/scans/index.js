Tea.context(function () {
  this.bScanning = false; //是否在扫描中
  this.checkValues = []; //选中的ID

  this.onStopScan = function () {
    if (this.checkValues.length > 0 && bScanning) {
      console.log("touch onStopScan");
    }
  };

  this.onCreateReport = function () {
    if (this.checkValues.length > 0) {
      console.log("touch onCreateReport");
    }
  };

  this.onDelete = function () {
    if (this.checkValues.length > 0 && !bScanning) {
      console.log("touch onDelete");
    }
  };

  this.clickCheckbox = function () {
    var checkDomArr = document.querySelectorAll(
      ".multi-table tbody input[type=checkbox]:checked"
    );
    this.checkValues = [];
    for (var i = 0, len = checkDomArr.length; i < len; i++) {
      this.checkValues.push(checkDomArr[i].value);
    }
    var allCheckDomArr = document.querySelectorAll(
      ".multi-table tbody input[type=checkbox]"
    );
    var allCheckbox = document.getElementById("js-all-checkbox");
    for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
      if (!allCheckDomArr[i].checked) {
        if (allCheckbox.checked) allCheckbox.checked = false;
        break;
      } else if (i === len - 1) {
        document.getElementById("js-all-checkbox").checked = true;
        break;
      }
    }
    this.updateBtnStatus();
  };
  this.checkAll = function () {
    var curClickBox = document.getElementById("js-all-checkbox")
    var allCheckDomArr = document.querySelectorAll(
      ".multi-table tbody input[type=checkbox]"
    );
    if (!curClickBox.checked) {
      // 点击的时候, 状态已经修改, 所以没选中的时候状态时true
      this.checkValues = [];
      for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
        var checkStatus = allCheckDomArr[i].checked;
        if (checkStatus) allCheckDomArr[i].checked = false;
      }
    } else {
      this.checkValues = [];
      for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
        var checkStatus = allCheckDomArr[i].checked;
        if (!checkStatus) allCheckDomArr[i].checked = true;
        this.checkValues.push(allCheckDomArr[i].value);
      }
    }
    this.updateBtnStatus();
  };
  this.updateBtnStatus = function () {
    const stopBtn = document.getElementById("stop-btn");
    const createBtn = document.getElementById("create-btn");
    const delBtn = document.getElementById("del-btn");
    if (this.checkValues.length > 0) {
      stopBtn.style.backgroundColor = "#14539A";
      stopBtn.style.cursor = "pointer";

      createBtn.style.backgroundColor = "#14539A";
      createBtn.style.cursor = "pointer";

      delBtn.style.backgroundColor = "#D9001B";
      delBtn.style.cursor = "pointer";
    } else {
      stopBtn.style.backgroundColor = "#AAAAAA";
      stopBtn.style.cursor = null;

      createBtn.style.backgroundColor = "#AAAAAA";
      createBtn.style.cursor = null;

      delBtn.style.backgroundColor = "#AAAAAA";
      delBtn.style.cursor = null;
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
