Tea.context(function () {
    this.keyword = ""
    this.curIndex = -1
  
    this.onGoBack = function () {
      window.location = "/hids/risk";
    }
  
    this.onOpenDetail = function (id) {
      window.location = "/hids/risk/weakList";
    }
  
    this.enters = function (index) {
      this.curIndex = index;
    }
  
    this.leaver = function (index) {
      this.curIndex = -1;
    }
  
    this.tableData = [
      {
        id: 1,
        hostData: {
          intIp: "192.168.0.1",
          outIp: "192.168.1.1",
          hostName: "elk.novalocal",
          systemName: "centos linux7.6.1810_64bit",
          rootName: "3.10.0-957.1.3.el7",
          macAddr: "3.10.0-957.1.3.el7",
          remarks: "备注信息",
        },
        value1: 10,
        value2: 11,
        value3: 3,
        value4: 5,
        value5: 8,
      },
      {
        id: 2,
        hostData: {
          intIp: "192.168.0.1",
          outIp: "192.168.1.1",
          hostName: "elk.novalocal",
          systemName: "centos linux7.6.1810_64bit",
          rootName: "3.10.0-957.1.3.el7",
          macAddr: "3.10.0-957.1.3.el7",
          remarks: "备注信息",
        },
        value1: 10,
        value2: 11,
        value3: 3,
        value4: 5,
        value5: 8,
      },
    ]
  })