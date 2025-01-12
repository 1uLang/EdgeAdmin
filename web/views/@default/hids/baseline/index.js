Tea.context(function () {
    this.localProgressData = localStorage.getItem("baselineProgressData")
    this.progressListData = this.localProgressData ? JSON.parse(this.localProgressData) : [];//{id:1,curPer:1,disabled:1}
    
    this.curIndex = -1
    this.sSelectValue = 0

    let that = this

    this.$delay(function () {

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
        // that.onCreateLoopTimeOut()
        window.addEventListener('beforeunload', function () {
            that.onReleaseUpdateTimeOut()
            // that.onReleaseTimeOut()
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
    
    this.getProgressItemInfo = function (id) {
        if(id){
            for(var index=0;index<that.progressListData.length;index++){
                if(that.progressListData[index].id == id){
                    return that.progressListData[index]
                }
            }
        }
        return null
    }

    this.getProgressPerStr = function (curValue,maxValue,id,state) { 
        console.log(curValue);
        if(!that.getProgressItemInfo){return "0%"}

        if(curValue == 0 ){
            if(state==1){
                that.onCreateUpdateTimeOut()
                let curData = that.getProgressItemInfo(id)
                if(curData){
                    if(curData.curPer == 0){
                        return "1%"
                    }
                    return curData.curPer+"%"
                }
                that.onCreateProgressItemInfo(id)
                return "1%"
            }else{
                that.onChangeProgressDataState(id,state)
                return ""
            }
        }else if(curValue == 100){
            that.onChangeProgressDataState(id,state)
        }

        if(curValue && maxValue && maxValue>0 && maxValue >= curValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            if(tempValue>=100){
                return "100%"
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
    
    this.getProgressPer = function (curValue, maxValue,id,state) {

        if(!that.getProgressItemInfo){return "0%"}

        if(curValue == 0 ){
            if(state && state==1){
                let curData = that.getProgressItemInfo(id)
                if(curData){
                    return curData.curPer+"%"
                }
                that.onCreateProgressItemInfo(id)
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

            //选择时间之后的回调
    this.onTimeChange = function () {
        let startTime = document.getElementById("day-from-picker").value
        if(startTime.length > 0){
            let tempCharCount=startTime.split(":").length-1;
            if(tempCharCount<=1){
                document.getElementById("day-from-picker").value = ""
            }else{
                startTime = startTime.replace("T", " ");
            }
        }else{
            document.getElementById("day-from-picker").value = ""
        }

        let endTime = document.getElementById("day-to-picker").value
        if(endTime.length > 0){
            let tempCharCount=endTime.split(":").length-1;
            if(tempCharCount<=1){
                document.getElementById("day-to-picker").value = ""
            }else{
                endTime = endTime.replace("T", " ");
            }
        }else{
            document.getElementById("day-to-picker").value = ""
        }
        //todo req
    }

    this.updateTimeId = null

    this.onCreateProgressItemInfo = function (id) {
        let curData = {id:id,curPer:0,state:1,disabled:0}
        that.progressListData.push(curData)
        that.onSaveProgressData()
    }

    this.onChangeProgressDataState = function (id,state) {
        for(var index=0;index<that.progressListData.length;index++){
            if(that.progressListData[index].id==id){
                that.progressListData[index].state = state
                break
            }
        }
        if(that.progressListData.length>0){
            that.progressListData = that.progressListData.filter((item) => {
                return item.state == 1;
            });
        }
        
        that.onSaveProgressData()
    }

    // 进度的缓存数据
    this.onUpdateProgressData = function () {
        for(var index=0;index<that.progressListData.length;index++){
            if(that.progressListData[index].state==1){
                that.progressListData[index].curPer = that.progressListData[index].curPer+5
                if(that.progressListData[index].curPer>=95){
                    that.progressListData[index].curPer = 95
                }
            }
        }
        that.onSaveProgressData()
    }

    this.onSaveProgressData = function () {
        localStorage.setItem("baselineProgressData", JSON.stringify(that.progressListData));
    }

    //计时器
    this.onCreateUpdateTimeOut = function () {
        if(!that.updateTimeId){
            that.updateTimeId = createTimer(that.onUpdateProgressData, {timeout: 5000});
            that.updateTimeId.start();
        }
    }

    this.onReleaseUpdateTimeOut = function () {

        if (that.updateTimeId) {
            that.updateTimeId.stop()
            that.updateTimeId = null
        }
        if(that.progressListData.length>0){
            that.progressListData = that.progressListData.filter((item) => {
                return item.state == 1;
            });
            that.onSaveProgressData()
        }
    }
});
  