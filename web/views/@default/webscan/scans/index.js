Tea.context(function () {
  this.bScanning = false; //是否在扫描中
  this.checkValues = []; //选中的ID

  this.nShowState = 1   //三个界面的状态控制 1 2 3

  this.bLoopholeDetail = false    //漏洞详情是否显示
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

    this.onChangeState = function (state) { 
        this.nShowState = state;
    }

    this.onShowDetail = function (id) { 
        this.onChangeState(2)
    }

    this.onOpenLoopholeDetail = function (id){
        if(id){
            bLoopholeDetail = true
        }else{
            bLoopholeDetail = false
        }
    }



    this.loopholeData = [
        {status:"height",desc:"zheshimiaoshu",address:"192.168.0.1"},
        {status:"low",desc:"zheshimiaoshu",address:"192.168.0.2"},
        {status:"middle",desc:"zheshimiaoshu",address:"192.168.0.3"},
    ]

    this.scans = [
       {id:1,address:"192.168.0.1",description:"sfjla",high:1,medium:2,low:3,info:4,status:"asdf",start_date:"2021-12-12T10:20:32.000"},
       {id:1,address:"192.168.0.1",description:"sfjla",high:1,medium:2,low:3,info:4,status:"asdf",start_date:"2021-12-12T10:20:32.000"},
       {id:1,address:"192.168.0.1",description:"sfjla",high:1,medium:2,low:3,info:4,status:"asdf",start_date:"2021-12-12T10:20:32.000"},
       {id:1,address:"192.168.0.1",description:"sfjla",high:1,medium:2,low:3,info:4,status:"asdf",start_date:"2021-12-12T10:20:32.000"},
    ]

    this.data = {
        scanTime:"13m32s", scanReq:"4,551",scansRet:"27ms",define:"123",
        attTag:[
            {id:1,key:"地址",value:"192.168.111.111",},
            {id:1,key:"地址",value:"192.168.111.111",},
            {id:1,key:"地址",value:"192.168.111.111",},
            {id:1,key:"地址",value:"192.168.111.111",},
            {id:1,key:"地址",value:"192.168.111.111",}
        ],
        warring:[
            {id:1,key:"Elasticsearch 服务可访问",value:"2021-12-12T10:20:32.000"}, 
            {id:1,key:"Elasticsearch 服务可访问",value:"2021-12-12T10:20:32.000"}, 
            {id:1,key:"Elasticsearch 服务可访问",value:"2021-12-12T10:20:32.000"}, 
            {id:1,key:"Elasticsearch 服务可访问",value:"2021-12-12T10:20:32.000"}, 
        ]
            
    }
});
