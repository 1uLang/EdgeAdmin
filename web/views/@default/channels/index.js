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
        window.location = "/users?selectChan=" + item.id
    }
    
    //替换参数
    this.onChangeItem = function(item){
        teaweb.popup(Tea.url(".create?id="+item.id), {
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
        teaweb.confirm("确定删除？",function() {
            this.$post(".delete").params({
                id: item.id,
            }).refresh()
        })
    }
    
    // this.tableData = [
    //     {id:1,name:"河北发发发公司",user:"张哥",mobile:"1300000000",productName:"发发发等保平台",clientCount:"3",status:1,createTime:"2021-09-13 08:22:22"},
    //     {id:1,name:"河北发发发公司",user:"张哥",mobile:"1300000000",productName:"发发发等保平台",clientCount:"3",status:0,createTime:"2021-09-13 08:22:22"},
    //     {id:1,name:"河北发发发公司",user:"张哥",mobile:"1300000000",productName:"发发发等保平台",clientCount:"3",status:1,createTime:"2021-09-13 08:22:22"},
    // ]
});
