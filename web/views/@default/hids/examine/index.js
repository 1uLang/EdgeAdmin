Tea.context(function () {

    this.nCheckState = 1    //体检状态
    this.nResultState = 1   //体检分数
    this.nNumState = 1      //分数类型
    this.sSelectValue = []  //体检项目
    this.dayFrom = ""
    this.dayTo = ""

    this.curIndex = -1

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })

})