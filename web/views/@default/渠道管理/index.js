Tea.context(function () {

    this.getStatus = function(status){
        switch(status){
            case 1:
                return "已启动"
            default:
                return "已停用"
        }
    }

    this.onCreateChannel = function(){
        teaweb.popup(Tea.url(".create"), {
            width:"600px",
            height:"570px",
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload();
                });
            },
        });
    }
    
    this.onLookDetail = function(item){

    }
    
    //替换参数
    this.onChangeItem = function(item){
        teaweb.popup(Tea.url(".create?xxxxx"), {
            width:"600px",
            height:"570px",
            callback: function () {
                teaweb.success("修改成功", function () {
                    teaweb.reload();
                });
            },
        });
    }

    this.onDeleteItem = function(item){

    }
    
    this.tableData = [
        {id:1,companyName:"河北发发发公司",linkPerson:"张哥",phoneNum:"1300000000",productName:"发发发等保平台",clientCount:"3",status:1,createTime:"2021-09-13 08:22:22"},
        {id:1,companyName:"河北发发发公司",linkPerson:"张哥",phoneNum:"1300000000",productName:"发发发等保平台",clientCount:"3",status:0,createTime:"2021-09-13 08:22:22"},
        {id:1,companyName:"河北发发发公司",linkPerson:"张哥",phoneNum:"1300000000",productName:"发发发等保平台",clientCount:"3",status:1,createTime:"2021-09-13 08:22:22"},
    ]
});
