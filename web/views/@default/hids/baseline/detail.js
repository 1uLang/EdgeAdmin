Tea.context(function () {

    this.nTableState = 1
    this.bShowDetail = false
    this.pDetailData = null

    this.$delay(function () {
	})

    this.onChangeTableState= function (state) {
        if(this.nTableState != state){
            this.nTableState = state
        }
      }


    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time && time.length > 0) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };



    this.onGoBack = function (id) { 
        window.location = "/hids/baseline";
    }
   
    this.getProgressPer = function (curValue,maxValue) { 
        if(curValue && maxValue && maxValue >= curValue){
            return (curValue/maxValue).toFixed(2)+"%"
        }
        return "0.00%"
     }
    
     this.onOpenDetail = function(id,tableData){
        this.getItemInfo(id,tableData)
        if(this.pDetailData){
            this.bShowDetail = true
        }
     }
     
     this.onCloseDetail = function(){
        this.bShowDetail = false
     }

     this.getItemInfo = function (id,tableData) { 
        this.pDetailData = null
        for (var i = 0; i < tableData.length; i++) {
            if (tableData[i].id == id) {
                this.pDetailData = tableData[i]
                break
            }
        }
      }

    this.valueData = {
        id:1,
        modelName:"linux_safedog_web_weblogic",
        checkCount:11,
        exceptionCount:0,
        checkTime:"2021-06-07T11:33:06.000",
        checkProject:"中间件安全",
        exceptionSystem:[
            {id:1,type:"安全审计",desc:"应启用安全审计功能，审计覆盖到每个用户，对重要的用户行为和重要安全事件进行审计",
            fixFunc:"在管理工具打开本地安全策略，打开路径:`安全设置\本地策略\审核策略`，将全部审核策略配置为：`成功,失败`。",
            checkFunc:"开启审核策略，对重要的用户行为和重要安全事件进行审计"}
        ],
        exceptionKey:[
            {id:1,type:"中间件安全1",desc:"审计覆盖到每个用户，对重要的用户行为和重要安全事件进行审计",
            fixFunc:"在管理工具打开本地安全策略，打开路径:`安全设置\本地策略\审核策略`，将全部审核策略配置为：`成功,失败`。",
            checkFunc:"开启审核策略，对重要的用户行为和重要安全事件进行审计"}
        ]
    }

});
  