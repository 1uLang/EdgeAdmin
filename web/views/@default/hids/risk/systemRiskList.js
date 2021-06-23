Tea.context(function () {
  this.nTableState = 1;
  this.dialogVisible = false;
  this.detailInfo=null

  this.onChangeTableState = function (state) {
    if (this.nTableState != state) {
      this.nTableState = state;
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

  this.getStatusName = function (status) {
    switch (status) {
      case 1:
        return "已修复";
      case 2:
        return "待修复";
      case 3:
        return "修复中";
      case 4:
        return "修复失败";
      case 5:
        return "已忽略";
      default:
        "未知";
    }
  };

  this.getDangerName = function (status) {
    switch (status) {
      case 1:
        return "低危";
      case 2:
        return "中危";
      case 3:
        return "高危";
      case 4:
        return "危机";
      default:
        "未评级";
    }
  };

  this.onGoBack = function () {
    window.location = "/hids/risk/systemRisk";
  };

  this.onDetail = function (id,tableData) {
    this.onGetItemInfo(id,tableData)
    if(this.detailInfo){
      this.dialogVisible = true;
    }
  };
  this.onFix = function (id) {};
  this.onIgnore = function (id) {};
  this.onCloseIgnore = function (id) {};

  this.onCloseDetail = function () {
    this.dialogVisible = false;
  };

  this.onGetItemInfo = function (id,tableData) { 
    this.detailInfo=null
    if(tableData){
      for (var i = 0; i < tableData.length; i++) {
        if (tableData[i].id == id) {
          this.detailInfo = tableData[i]
          break
        }
      }
    }
    
  }

  this.dataInfo = {
    systemInfo: "centos linux7.6.1810_64bit",
    lastScanTime: "2021-06-07T11:33:06.000",
    loopholeCount: 23,
  };

  this.tableData1 = [
    {
      id: 1,
      level: 1,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 1,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    },
    {
      id: 2,
      level: 2,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 2,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    },
    {
      id: 3,
      level: 3,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 3,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    },
    {
      id: 4,
      level: 4,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 4,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    },
    {
      id: 5,
      level: 4,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 5,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    },
  ];
  this.tableData2 = [
    {
      id: 1,
      level: 1,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 1,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    },
    {
      id: 2,
      level: 2,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 2,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    },
    {
      id: 3,
      level: 3,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 3,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    },
    {
      id: 4,
      level: 4,
      value1: "KB4493475 2019 年 4 月 9 日 - KB44934",
      value2: "WEB容器配置缺陷",
      value3: "需重启系统",
      value4: "2021-06-07T11:33:06.000",
      value5: "本地扫描",
      value6: 4,
      detailTime: "2021-06-07T11:33:06.000",
      systemVer:"10240.18186",
      dangerSource:"Microsoft",
      dangerInfo:"此更新包修改包括：1.解决了在节点操作期间引发异常时，可能导致使用 MSXML6 的应用程序停止响应的问题。2.解决了在编辑包含 Internet Explorer 10 Internet 设置的组策略首选项 (GPP) 的组策略对象 (GPO) 时，导致组策略编辑器停止响应的问题。",
      url:"https://support.microsoft.com/en-us/kb/4493475"
    }
  ];
});
