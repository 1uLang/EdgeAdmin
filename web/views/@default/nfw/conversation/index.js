Tea.context(function () {
    // this.keyword = ""
    //directionState 方向 1 进 2 出
    this.getDirection=function (state) { 
        switch(state){
            case "in":
                return "进"
            case "out":
                return "出"
            default:
                return "未知"
        }
    }

    this.getAddr = function (ip,port) {
        if(ip != null){
            return ip+":"+port
        }
        return ""
    }

    //
    this.getNumber = function (num) {
        if(num < 1024){
            return num
        }
        return Math.floor(num /1024) + " K"
    }

    this.onSearch = function () {
        window.location.href = '/nfw/conversation?nodeId=' + this.selectNode+"&keyword="+this.keyword
    }
    //获取当前选中的节点
    this.GetSelectNode = function (event) {
        this.selectNode = event.target.value; //获取option对应的value值
        // localStorage.setItem("nfwSelectNodeId", this.selectNode);
        let node = this.selectNode
        window.location.href = '/nfw/conversation?nodeId=' + node+"&keyword="+this.keyword

    }

    // this.tableData = [
    //     {id:1,directionState:1,httpType:"udp",source:"192.168.101.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:1,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"},
    //     {id:2,directionState:0,httpType:"tcp",source:"192.168.102.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:1,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"},
    //     {id:3,directionState:1,httpType:"udp",source:"192.168.103.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:2,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"},
    //     {id:4,directionState:0,httpType:"tcp",source:"192.168.104.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:1,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"},
    //     {id:5,directionState:1,httpType:"udp",source:"192.168.105.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:1,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"}
    // ]
})