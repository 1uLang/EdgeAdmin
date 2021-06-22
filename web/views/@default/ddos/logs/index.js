Tea.context(function () {

    this.links = []
    this.traffics = []
    this.level = 1
    this.address = ''
    this.nShowState = 1

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })


    this.showHost = function () { //重新加载该页面
        let node = this.nodeId
        window.location.href = '/ddos/logs?nodeId=' + node
    }
    this.onDelete = function (id) {

    }

    this.onChangeShowState = function (state) {
        if (this.nShowState != state) {
            if (state === 2) {
                this.$get(".traffic").params({NodeId: this.nodeId}).success(resp => {
                    if (resp.code === 200) {
                        if (resp.data.traffics)
                            this.traffics = resp.data.traffics
                        else
                            this.traffics = []
                        this.level = resp.data.level
                        this.nShowState = state
                    }
                })
            } else {
                this.$get(".link").params({NodeId: this.nodeId}).success(resp => {
                    if (resp.code === 200) {
                        if (resp.data.links)
                            this.links = resp.data.links
                        else
                            this.links = []
                        this.level = resp.data.level
                        this.nShowState = state
                    }
                })
            }

        }
    }
    this.search = function () {
        if(this.nShowState ==2){
            this.$get(".traffic").params({NodeId: this.nodeId,Address:this.address,Level:this.level}).success(resp => {
                if (resp.code === 200) {
                    if (resp.data.traffics)
                        this.traffics = resp.data.traffics
                    else
                        this.traffics = []
                    this.level = resp.data.level
                }
            })
        }else if(this.nShowState == 3){
            this.$get(".link").params({NodeId: this.nodeId,Address:this.address,Level:this.level}).success(resp => {
                if (resp.code === 200) {
                    if (resp.data.links)
                        this.links = resp.data.links
                    else
                        this.links = []
                    this.level = resp.data.level
                }
            })
        }
    }
    this.toShowStatus = function (st) {
        if (st === "2") {
            return "结束"
        } else {
            return "保护中"
        }
    };

    this.onDownLoad = function (id) {

    }

})