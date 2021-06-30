Tea.context(function () {
    this.sqlName = ""

    this.getStatus = function (status) {
        switch (status) {
            case '1':
                return "启已用"
            case '0':
                return "已停用"
            default:
                return "已停用"
        }
    }

    this.onAuth = function (id) { 

    }
    this.onEnabled = function (id) { 

    }
    this.onDisabled =function (id) {

    }
    this.onDelete = function (id) { 
        
     }
 
    this.tableData = [
        {id:1,value1:"robin_mysql",value2:"47.108.234.195",value3:"3306",value4:"mysql",value5:"5.8",value6:"linux",status:1},
        {id:2,value1:"robin_mysql",value2:"47.108.234.195",value3:"3306",value4:"mysql",value5:"5.8",value6:"linux",status:0},
        {id:3,value1:"robin_mysql",value2:"47.108.234.195",value3:"3306",value4:"mysql",value5:"5.8",value6:"linux",status:1},
        {id:4,value1:"robin_mysql",value2:"47.108.234.195",value3:"3306",value4:"mysql",value5:"5.8",value6:"linux",status:0}
    ]
})