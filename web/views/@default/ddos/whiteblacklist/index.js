Tea.context(function () {
    this.severity=''
    this.onAddNameList = function () { 
        teaweb.popup(Tea.url(".createPopup"), {
            callback: function () {
              teaweb.success("保存成功", function () {
                teaweb.reload();
              });
            },
          });
    }

    this.onDelete = function (id) { 

    }

    this.tableData = [
        {id:1,host:"192.168.0.1",nameType:"黑名单",matchCount:303,remarks:"备注"},
        {id:1,host:"192.168.0.1",nameType:"黑名单",matchCount:303,remarks:"备注"},
        {id:1,host:"192.168.0.1",nameType:"黑名单",matchCount:303,remarks:"备注"},
        {id:1,host:"192.168.0.1",nameType:"黑名单",matchCount:303,remarks:"备注"},
    ]
})