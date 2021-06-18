Tea.context(function () {
  this.checkValues = []; //选中的ID

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
    const delBtn = document.getElementById("del-btn");
    if (this.checkValues.length > 0) {
      delBtn.style.backgroundColor = "#D9001B";
      delBtn.style.cursor = "pointer";
    } else {
      delBtn.style.backgroundColor = "#AAAAAA";
      delBtn.style.cursor = null;
    }
  };

  //item 内的删除
  this.deleteNode = function (id) {
    teaweb.confirm("确定要删除这个节点吗？", function () {});
  };
  //btn删除
  this.onDelete = function () {
    teaweb.confirm("确定要删除这些节点吗？", function () {});
  };

  this.onChangeTimeFormat = function (time) {
    var resultTime = "";
    if (time) {
      var tempTime = time.substring(0, time.indexOf("."));
      resultTime = tempTime.replace("T", " ");
    }
    return resultTime;
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
