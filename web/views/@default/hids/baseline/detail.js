Tea.context(function () {

    this.nTableState = 1
    this.bShowDetail = false
    this.pDetailData = {}

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

     this.onCloseDetail = function(){
        this.bShowDetail = false
         this.pDetailData = {}
     }

     this.onOpenDetail = function (suggest,checkFunc) {
         this.pDetailData.suggest = suggest
         this.pDetailData.checkFunc = checkFunc
         this.bShowDetail = true
      }
});
  