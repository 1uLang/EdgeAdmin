Tea.context(function () {
    this.severity=''

    this.onDelete = function (id) { 

    }

    this.nShowState = 1

    this.onChangeShowState = function (state){
        if(this.nShowState != state){
            this.nShowState = state
        }
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time && time.length > 0) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }else{
            resultTime = "— —"
        }
        return resultTime;
    };

    this.onDownLoad = function (id) { 
        
     }


    this.tableData1 = [
        {id:1,tagAddress:"118.112.250.249",value1:"tcp 80",value2:"2021-06-06T19:33:41.000",value3:"2021-06-06T19:33:41.000",value4:"[SYN]",value5:"[SlowHTTP]",value6:"结束",value7:"2457.58",value8:"223.104.113.0/24"},
        {id:1,tagAddress:"118.112.250.249",value1:"tcp 80",value2:"2021-06-06T19:33:41.000",value3:"2021-06-06T19:33:41.000",value4:"[SYN]",value5:"[SlowHTTP]",value6:"结束",value7:"2457.58",value8:"223.104.113.0/24"},
        {id:1,tagAddress:"118.112.250.249",value1:"tcp 80",value2:"2021-06-06T19:33:41.000",value3:"",value4:"[SYN]",value5:"[SlowHTTP]",value6:"结束",value7:"2457.58",value8:"223.104.113.0/24"},
        {id:1,tagAddress:"118.112.250.249",value1:"tcp 80",value2:"2021-06-06T19:33:41.000",value3:"2021-06-06T19:33:41.000",value4:"[SYN]",value5:"[SlowHTTP]",value6:"结束",value7:"2457.58",value8:"223.104.113.0/24"},
    ]

    this.tableData2 = [
        {id:1,censusData:"2021-06-05",value1:"20:20",value2:32534,value3:14056,value4:14056,value5:14056,value6:14056,value7:14056,value8:14056,value9:14056,value10:14056,value11:14056},
        {id:1,censusData:"2021-06-05",value1:"20:20",value2:32534,value3:14056,value4:14056,value5:14056,value6:14056,value7:14056,value8:14056,value9:14056,value10:14056,value11:14056},
        {id:1,censusData:"2021-06-05",value1:"20:20",value2:32534,value3:14056,value4:14056,value5:14056,value6:14056,value7:14056,value8:14056,value9:14056,value10:14056,value11:14056},
        {id:1,censusData:"2021-06-05",value1:"20:20",value2:32534,value3:14056,value4:14056,value5:14056,value6:14056,value7:14056,value8:14056,value9:14056,value10:14056,value11:14056},
    ]

    this.tableData3 = [
        {id:1,censusData:"2021-06-05",value1:"20:20",value2:32534,value3:14056,value4:14056,value5:14056,value6:14056,value7:14056,value8:14056,value9:14056},
        {id:1,censusData:"2021-06-05",value1:"20:20",value2:32534,value3:14056,value4:14056,value5:14056,value6:14056,value7:14056,value8:14056,value9:14056},
        {id:1,censusData:"2021-06-05",value1:"20:20",value2:32534,value3:14056,value4:14056,value5:14056,value6:14056,value7:14056,value8:14056,value9:14056},
        {id:1,censusData:"2021-06-05",value1:"20:20",value2:32534,value3:14056,value4:14056,value5:14056,value6:14056,value7:14056,value8:14056,value9:14056},
    ]
})