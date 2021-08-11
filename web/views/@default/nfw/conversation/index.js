Tea.context(function () {
    this.keyword = ""
    //directionState 方向 1 进 2 出
    this.getDirection=function (state) { 
        switch(state){
            case 1:
                return "进"
            case 2:
                return "出"
            default:
                return "出"
        }
    }

    this.getTagState = function (state) { 
        switch(state){
            case 1:
                return "success"
            case 2:
                return "fail"
        }
    }

    this.onSearch = function () { 

    }
    this.tableData = [ 
        {id:1,directionState:1,httpType:"udp",source:"192.168.101.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:1,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"},
        {id:2,directionState:0,httpType:"tcp",source:"192.168.102.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:1,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"},
        {id:3,directionState:1,httpType:"udp",source:"192.168.103.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:2,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"},
        {id:4,directionState:0,httpType:"tcp",source:"192.168.104.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:1,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"},
        {id:5,directionState:1,httpType:"udp",source:"192.168.105.1:4404",gateway:"182.168.101.126:4559",tagAddr:"182.168.101.126:4559",state:1,ageSec:"23545",expires:"86544",dataSize:"39.55K",byteSize:"3.18M"}
    ]
})