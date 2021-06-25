Tea.context(function () {
  this.checkValues = []; //选中的ID
  this.checkTargetValues = []; //选中的targetID

  this.nShowState = 1   //三个界面的状态控制 1 2 3
  this.vulnerabilities = []
  this.statistics = {}
  this.checkPer = "12%"
  this.Address = ""
  this.bLoopholeDetail = false    //漏洞详情是否显示
  this.onStopScan = function () {
    if (this.checkValues.length > 0) {
      let that = this
      let scan_ids = JSON.parse(JSON.stringify(this.checkValues))
      teaweb.confirm("确定要停止这个扫描吗？", function () {
        that.$post(".stop")
            .params({
              Ids: scan_ids
            })
            .refresh()
      })
    }
  };

  this.onCreateReport = function () {
    if (this.checkValues.length > 0) {
      let that = this
      let scan_ids = JSON.parse(JSON.stringify(this.checkValues))
      let tarId = JSON.parse(JSON.stringify(this.checkTargetValues))
      teaweb.confirm("确定要生成这个扫描的报表吗？", function () {
        that.$post("/web/scan/report/create")
            .params({
              Ids: scan_ids,
              TarIds: tarId,
            })
            .refresh()
      })
    }
  };

  this.onDelete = function () {
    if (this.checkValues.length > 0 ) {
      let that = this
      let scan_ids = JSON.parse(JSON.stringify(this.checkValues))
      teaweb.confirm("确定要删除这个扫描吗？", function () {
        that.$post(".delete")
            .params({
              ScanIds: scan_ids
            })
            .refresh()
      })
    }
  };

  this.clickCheckbox = function () {
    var checkDomArr = document.querySelectorAll(
      ".multi-table tbody input[type=checkbox]:checked"
    );
    this.checkValues = [];
    this.checkTargetValues = [];
    for (var i = 0, len = checkDomArr.length; i < len; i++) {
      this.checkValues.push(checkDomArr[i].value);
      let tar  = checkDomArr[i].getAttribute("data")
      console.log("1targetid="+tar);
      this.checkTargetValues.push(tar);
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
      this.checkTargetValues = [];
      for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
        var checkStatus = allCheckDomArr[i].checked;
        if (checkStatus) allCheckDomArr[i].checked = false;
      }
    } else {
      this.checkValues = [];
      this.checkTargetValues = [];
      for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
        var checkStatus = allCheckDomArr[i].checked;
        if (!checkStatus) allCheckDomArr[i].checked = true;
        this.checkValues.push(allCheckDomArr[i].value);
        let tar  = allCheckDomArr[i].getAttribute("data")
        console.log("2targetid="+tar);
        this.checkTargetValues.push(tar);
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
        if(state ===3){ //漏洞详情
          this.vulnerabilities = []
          this.$get("/webscan/vulnerabilities").params({Address:this.Address,List:true}).success(resp => {
            if(resp.code===200){
              this.vulnerabilities = resp.data.vulnerabilities
            }
          })
        }
    }

    this.onShowDetail = function (item) {

      //获取漏洞详情报表
      this.$get(".statistics").params({ScanId: item.scan_id,ScanSessionId:item.current_session.scan_session_id}).success(resp => {
        if(resp.code===200){
          console.log(item.target_id)
          this.statistics.status = this.onChangeStatusFormat(resp.data.statistics.status)
          this.statistics.severity_counts = item.current_session.severity_counts
          this.statistics.event_level = resp.data.statistics.scanning_app.wvs.event_level
          this.statistics.host = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].host
          this.statistics.os = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].target_info.os
          this.statistics.responsive = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].target_info.responsive ? '是':'否'
          this.statistics.server = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].target_info.server
          this.statistics.technologies = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].target_info.technologies
          this.statistics.request_count = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].web_scan_status.request_count
          this.statistics.avg_response_time = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].web_scan_status.avg_response_time
          this.statistics.locations = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].web_scan_status.locations
          this.statistics.vulns = resp.data.statistics.scanning_app.wvs.main.vulns
          this.statistics.duration = resp.data.statistics.scanning_app.wvs.main.duration
          this.statistics.progress = resp.data.statistics.scanning_app.wvs.main.progress
          this.statistics.messages = resp.data.statistics.scanning_app.wvs.main.messages
          console.log(this.statistics)
          this.onChangeState(2)
        }
      })
    }

    this.refreshProgress = function () { 
        var maxCount = 100
        var tempCount = 25
        var curPer = Math.floor(maxCount/tempCount)
        checkPer = curPer+"%"
        document.getElementById("barContent").style.width = checkPer;
     }

 
  this.onChangeStatusFormat = function (status) {
    var resultStatus = status;
    if (status) {
      switch (status) {
        case "aborted":return "已中止";
        case "completed": return "已完成";
        case "progressing": return "正在进行";
      }
    }
    return resultStatus;
  };
  this.onChangeTimeFormat2 = function (time){
    let m = parseInt(time / 60)
    let s = time % 60
    return m +'m' +s +'s'
  };
  this.onChangeCountFormat = function (count){
    let q = parseInt(count / 1000);
    let s = count % 1000;

    return q +','+s
  };
});
