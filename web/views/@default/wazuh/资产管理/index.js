Tea.context(function () {
    
    this.pageState = 1

    this.getLinkState=function(status){
        switch (status) {
            case 1:
                return "已连接"
            case 2:
                return "未连接"
            default:
                return "未知"
        }
    }
    
    this.onAddHost = function(){
        teaweb.popup(Tea.url(".create"), 
            {
                width:"650px",
                height:"400px",
                callback: function () {
                    teaweb.success("保存成功", function () {
                        teaweb.reload();
                    });
                },
            },
        );
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.onDelete = function(item){
        teaweb.confirm("确定要删除？", function () {
            
        })
    }

    this.tableData = [
        {
            id:1,
            hostIP:"192.168.0.1",
            hostName:"WIN-OC98D604V3R1",
            systemName:"CentOS Linux7.6",
            createTime:"2021-02-02T15:25:30.123",
            lastTime:"2021-02-02T15:25:30.123",
            status:2,
        },
        {
            id:2,
            hostIP:"192.168.0.2",
            hostName:"WIN-OC98D604V3R2",
            systemName:" Microsoft Windows 10 中国家庭版 10.0.19042",
            createTime:"2021-02-02T15:25:30.123",
            lastTime:"2021-02-02T15:25:30.123",
            status:1,
        },
        {
            id:3,
            hostIP:"192.168.0.3",
            hostName:"WIN-OC98D604V3R3",
            systemName:" Microsoft Windows 10 中国家庭版 10.0.19042",
            createTime:"2021-02-02T15:25:30.123",
            lastTime:"2021-02-02T15:25:30.123",
            status:0,
        },
    ]
})