Tea.context(function () {
    this.curIndex = -1
    this.sSelectValue = 0

    this.dayFrom = ""
    this.dayTo = ""

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })
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
                return "未检查"
        }
     }

     this.enters = function (index) {
        // this.curIndex = index;
      }
    
      this.leaver = function (index) {
        this.curIndex = -1;
      }
    
    this.getProgressPerStr = function (curValue,maxValue) { 
    if(curValue && maxValue && maxValue >= curValue){
        let value = parseInt(curValue/maxValue * 100)
        if(value>=100){
            return "已完成"
        }
        return value+"%"
    }
    return "1%"
    }

    this.getProgressPer = function (curValue,maxValue) { 
        if(curValue && maxValue && maxValue >= curValue){
            let value = parseInt(curValue/maxValue * 100)
            return value+"%"
        }
        return "1%"
     }

    //合规基线
     this.onOpenCheck = function (item) {
        //打开合规基线弹窗
         teaweb.popup(Tea.url(".template?macCode="+item.macCode+"&os="+item.os.osType), {
             height: "500px",
         })
      }

    this.onStartCheck = function () {
        teaweb.confirm("确定开始体检该主机吗？", function () {
            this.$post(".check").params({
                Opt: "ignore",
                MacCode: [this.macCode],
                templateId: this.sSelectValue,
            }).refresh()
        })
    }

    this.onOpenDetail = function (item) {
        window.location = "/hids/baseline/detail?macCode="+item.macCode+'&pageSize='+item.totalItemCount
    }

    //添加/删除元素
    this.onAddSelectValue = function (index) {
        console.log("=========",index)
        this.sSelectValue = index
    }
    this.getShowSelectImage = function (id) {
        console.log(id,this.sSelectValue)

        if (this.sSelectValue == id) {
          return "/images/select_select.png";
        }
        return "/images/select_box.png";
      }

      this.onTimeChange = function () {
      
      }
});
  