Tea.context(function () {

    this.bOpenRelease = false

    this.agentUrl = ""
    this.agentId = ""
    this.agentKey = ""
    this.jumpUrl = ""

    this.getStatus = function (status) {
        if (status) return "已启用"
        return "已停用"      
    }
    this.getLinkStatus = function(status){
        if (status) return "已连接"
        return "未连接"
    }

    this.onAddHost = function () {
        teaweb.popup(Tea.url(".createPopup"), {
            height: "310px",
            width: "500px",
        })
    }

    this.onEdit = function(item){
        let tempEnabled = 0
        if(item.on){
            tempEnabled = 1
        }
        teaweb.popup(Tea.url(".createPopup?id="+item.id+"&name="+escape(item.name)+"&address="+item.host+"&enabled="+tempEnabled+"&selectType="+item.os_type), {
            height: "310px",
            width: "500px",
        })
    }
    this.onRelease = function(item){
        // this.bOpenRelease = true
        this.$get("/resmon/agents").params(
            {id:item.id}
        ).success(function(res){
            this.agentUrl = res.data.dwon_info.host
            this.agentId = item.id
            this.agentKey = item.key
            this.jumpUrl = res.data.dwon_info.down_url
            this.bOpenRelease = true
        })   
    }
    this.onDelete = function(id){
        teaweb.confirm("确定要删除此主机吗？", function () {
            this.$post("/resmon/delete").params(
                {
                    id:id,
                }
            ).success(
                function(){
                    teaweb.success("删除成功", function () {
                        teaweb.reload()
                    })
                }
            )
        })
    }
    this.onJumpURL = function(){
        window.location = this.jumpUrl
    }
    
    this.onCloseRelease = function(){
        this.bOpenRelease = false
    }

    this.dbList = [
        {id:1,hostName:"审计系统",hostAddress:"192.168.1.1",hostPlatform:"Linux 64位",CPU:"5%",RAM:"0.6G/4G",disk:"12.7G/44.1G",status:1,enabled:1},
        {id:1,hostName:"审计系统",hostAddress:"192.168.1.1",hostPlatform:"Linux 64位",CPU:"5%",RAM:"0.6G/4G",disk:"12.7G/44.1G",status:1,enabled:1},
        {id:1,hostName:"审计系统",hostAddress:"192.168.1.1",hostPlatform:"Linux 64位",CPU:"5%",RAM:"0.6G/4G",disk:"12.7G/44.1G",status:1,enabled:1},
    ]
})