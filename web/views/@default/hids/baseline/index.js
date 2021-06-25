Tea.context(function () {
    this.curIndex = -1
    this.sSelectValue = []

    this.onChangeCheckState = function (state) {
        window.location = "/hids/baseline?State="+state+'&ResultState='+this.ResultState
     }

    this.onChangeResultState = function(state){
        window.location = "/hids/baseline?State="+this.State+'&ResultState='+state
    }
    this.parseServerLocalIp = function (ip){
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }

    this.getStateName = function (status) {
        switch(status){
            case 0:
                return "未检查"
            case 1:
                return "检查中"
            case 2:
                return "已完成"
            case 3:
                return "检查失败"
            default:
                return "未知"
        }
     }

     this.enters = function (index) {
        this.curIndex = index;
      }
    
      this.leaver = function (index) {
        this.curIndex = -1;
      }
    
    this.getProgressPer = function (curValue,maxValue) { 
        if(curValue && maxValue && maxValue >= curValue){
            return parseInt(curValue/maxValue * 100)+"%"
        }
        return "0%"
     }

    //合规基线
     this.onOpenCheck = function (item) {
        //打开合规基线弹窗
         teaweb.popup(Tea.url(".template?macCode="+item.macCode), {
             height: "30em",
         })
      }

    this.onStartCheck = function (macCode) {
        console.log(macCode)
    }

    this.onStopCheck = function (id) { 

    }

    this.onOpenDetail = function (item) {
        window.location = "/hids/baseline/detail?macCode="+item.macCode+'&pageSize='+item.totalItemCount
    }

    //检测是否包含元素
    this.checkSelectValue = function (index) {
        for (var i = 0; i < this.sSelectValue.length; i++) {
            if (this.sSelectValue[i] == index) {
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
            this.sSelectValue = this.sSelectValue.filter((itemIndex) => {
                return itemIndex != index;
            });
        } else {
            this.sSelectValue.push(index);
        }
    }
    this.getShowSelectImage = function (id) {
        let bValue = false;
        if(this.sSelectValue){
          bValue = this.checkSelectValue(id);
        }
        if (bValue) {
          return "/images/select_select.png";
        }
        return "/images/select_box.png";
      }
});
  