Tea.context(function () {
    this.curIndex = -1
    this.sSelectValue = 0

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
         teaweb.popup(Tea.url(".template?macCode="+item.macCode+"&os="+item.os.osType), {
             height: "30em",
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
});
  