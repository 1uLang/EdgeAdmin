Tea.context(function () {
  this.keyword = ""
  this.moreSelect = ""
  this.curIndex = -1
  this.showSelectValue = ""
  this.selectValue = []

  this.bShowSelectBox = false

  this.onGoBack = function () {
    window.location = "/hids/risk";
  }

  this.onOpenDetail = function (id) {
    window.location = "/hids/risk/systemRiskList";
  }

  this.openSelect = function () {
    this.bShowSelectBox = !this.bShowSelectBox;
  }

  this.closeSelect = function () {
    this.bShowSelectBox = false;
  }

  this.enters = function (index) {
    this.curIndex = index;
  }

  this.leaver = function (index) {
    this.curIndex = -1;
  }

  //检测是否包含元素
  this.checkSelectValue = function (index) {
    for (var i = 0; i < this.selectValue.length; i++) {
      if (this.selectValue[i] == index) {
        return true;
      }
    }
    return false;
  }

  //添加/删除元素
  this.onAddSelectValue = function (index) {
    let bValue = false;
    if(this.checkSelectValue){
      bValue = this.checkSelectValue(index);
    }
    if (bValue) {
      this.selectValue = this.selectValue.filter((itemIndex) => {
        return itemIndex != index;
      });
    } else {
      this.selectValue.push(index);
    }
    this.getShowSelectValue();
  }

  //显示选择的值
  this.getShowSelectValue = function () {
    this.showSelectValue = "";
    for (var i = 0; i < this.selectValue.length; i++) {
      for (var j = 0; j < this.dangerData.length; j++) {
        if (this.selectValue[i] == this.dangerData[j].id) {
          this.showSelectValue =
            this.showSelectValue + " " + this.dangerData[j].value;
          break;
        }
      }
    }
  }

  this.getShowSelectImage = function (id) {
    let bValue = false;
    if(this.checkSelectValue){
      bValue = this.checkSelectValue(id);
    }
    if (bValue) {
      return "/images/select_select.png";
    }
    return "/images/select_box.png";
  }

  this.dangerData = [
    { id: 1, value: "低危" },
    { id: 2, value: "中危" },
    { id: 3, value: "高危" },
    { id: 4, value: "危急" },
    { id: 5, value: "未评级" },
  ];

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