Tea.context(function () {

    this.nShowState = 1
    this.nTimeSelect = 1
    this.dayFrom = ""
    this.dayTo = ""

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })

    this.onChangeShowState = function (state){
        if(this.nShowState!=state){
            this.nShowState = state
        }
    }

    this.onChangeTimeSelect = function(index){
        if(this.nTimeSelect!=index){
            this.nTimeSelect = index
        }
    }
    this.data = {
        hostCount:5,
        hostOnline:1,
        detailAttCount:10,
        detailAttTodayCount:1,
        detailloopholeCount:120,
        detailloopholeTodayCount:12,
    }

});
  