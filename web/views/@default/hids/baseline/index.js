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
        let that = this
        that.onCreateLoopTimeOut()
        window.addEventListener('beforeunload', function () {
            that.onReleaseTimeOut()
        })
    })

    this.onCallBack = function () {
        if (this.checkScans()) {
            this.$post(".").success(resp => {
                if (resp.code === 200) {
                    this.baselines = resp.data.baselines
                    this.State = resp.data.State
                    this.ResultState = resp.data.ResultState
                }
            })
        }
    }
    this.onCreateLoopTimeOut = function () {
        this.onReleaseTimeOut()
        this.checkTimer = createTimer(this.onCallBack, {timeout: 60000});
        this.checkTimer.start();
    }
    this.onReleaseTimeOut = function () {
        if (this.checkTimer) {
            this.checkTimer.stop()
            this.checkTimer = null
        }
    }
    this.checkScans = function () {
        for (item of this.datas) {
            if (item.state === 1) {
                return true
            }
        }
        return false
    }

    this.onChangeCheckState = function (state) {
        window.location.href = "/hids/baseline?State="+state+'&ResultState='+this.ResultState
     }

    this.onChangeResultState = function(state){
        window.location.href = "/hids/baseline?State="+this.State+'&ResultState='+state
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
    
    this.getProgressPerStr = function (curValue,maxValue,state) { 
        if(curValue == 0 ){
            if(state==1){
                return "1%"
            }else if(state==0){
                return ""
            }
        }

        if(curValue && maxValue && maxValue>0 && maxValue >= curValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            if(tempValue>=100){
                return "已完成"
            }else if(tempValue<1 && state && state==1){
                return "1%"
            }
            
            return tempValue + "%"
        }
        return "0%"
    }

    this.checkShowColor = function (curValue, maxValue) {
        if(curValue && maxValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            return tempValue >= 100
        }
        return false
    }
    

    this.getProgressPer = function (curValue,maxValue,state) { 
        if(curValue == 0 ){
            if(state && state==1){
                return "1%"
            }
        }

        if(curValue && maxValue && maxValue>0 && maxValue >= curValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            if(tempValue<1 && state && state==1 ){
                return "1%"
            }
            return tempValue + "%"
        }
        return "0%"
     }

    //合规基线
     this.onOpenCheck = function (item) {

         let serverIp = item.serverIp

        //打开合规基线弹窗
         teaweb.popup(Tea.url(".template?macCode="+item.macCode+'&serverIp='+serverIp+"&os="+item.os.osType), {
             height: "500px",
         })
      }

    this.onOpenDetail = function (item) {
        this.serverIp = item.serverIp
        window.location.href = "/hids/baseline/detail?macCode="+item.macCode+'&pageSize='+item.totalItemCount+'&time='+item.overTime+'&checkCount='+item.riskItemCount
    }

    //添加/删除元素
    this.onAddSelectValue = function (index) {
        this.sSelectValue = index
    }
    this.getShowSelectImage = function (id) {
        if (this.sSelectValue == id) {
          return "/images/select_select.png";
        }
        return "/images/select_box.png";
      }

      this.onTimeChange = function () {
      
      }
});
  