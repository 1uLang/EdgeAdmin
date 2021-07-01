Tea.context(function () {
    
    this.dayFrom = ""
    this.dayTo = ""
    this.userName = ""
    this.systemUserName = ""
    this.cdmName = ""

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
          var tempTime = time.substring(0, time.indexOf("."));
          resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.onCheck = function(){

    }
    this.onReset = function () { 
        
     }


    this.tableData = [
        {   id:1,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd",
            value2:"普通",value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"dby_web",value6:"2021-06-23T16:00:00.125"
        },
        {   id:2,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd",
            value2:"普通",value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"dby_web",value6:"2021-06-23T16:00:00.125"
        },
        {   id:3,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd",
            value2:"普通",value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"dby_web",value6:"2021-06-23T16:00:00.125"
        },
        {   id:4,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd",
            value2:"普通",value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"dby_web",value6:"2021-06-23T16:00:00.125"
        },
    ]
});
  