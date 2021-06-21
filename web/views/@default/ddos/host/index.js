Tea.context(function () {

    //获取url上参数
    this.nodeId = function (variable) {
        var query = window.location.search.substring(1);
        var vars = query.split("&");
        for (var i = 0; i < vars.length; i++) {
            var pair = vars[i].split("=");
            if (pair[0] == variable) {
                return pair[1];
            }
        }
        return ('');
    }('nodeId')

    this.level = "0"//防护策略等级
    this.ignore = false //ip直通
    this.dropNode = ''
    this.address = ''
    this.src_ip = ""
    this.shieldList = []
    this.linkList = []
    this.nShowState = 1 //当前显示的页面

    this.tableState = 1 //只有在nShowState==2时生效 屏蔽和连接列表的切换

    this.onAddNodeIP = function () {
        let node = this.getNodeId()
        teaweb.popup(Tea.url(".createPopup?nodeId="+node), {
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload();
                });
            },
        });
    }
    this.showHost = function () { //重新加载该页面
        let node = ''
        if (this.nodeId === '') {    //重新加载该页面
            node = document.getElementById('selectBox').value
            this.nodeId = node
        } else {
            node = this.nodeId
        }
        window.location.href = '/ddos/host?nodeId=' + node
    }
    this.setHost = function (notice) {

        let ignore = document.getElementById('btn-switch-ignore').checked
        let level = document.getElementById('ddosLevel').value
        let that = this
        let msg = !this.ignore ? '开启' : '关闭'
        let node = this.getNodeId()
        if (notice) {
            teaweb.confirm("确定要" + msg + "IP直通吗？", function () {
                that.$post(".set").params({
                    Addr: that.address,
                    ignore: ignore,
                    NodeId: node,
                    set: level,
                }).refresh()
                document.getElementById("switch-btn").checked = !ignore
            })
        } else {
            if (this.level === level) {
                return
            }
            that.$post(".set")
                .params({
                    Addr: that.address,
                    ignore: ignore,
                    set: level,
                    NodeId: node,
                }).refresh()
        }
    }
    //分页切换
    this.changeState = function (state) {
        if (this.nShowState != state) {
            this.nShowState = state
        }
    }
    this.getNodeId = function () {
        let node = ''
        if (this.nodeId === '') {    //重新加载该页面
            node = document.getElementById('selectBox').value
            this.nodeId = node
        } else {
            node = this.nodeId
        }
        return node
    }
    this.shieldSearchList = function (state) {
        let searchIp = this.src_ip === '' ? '' : '-' + this.src_ip
        let node = this.getNodeId()
        //屏蔽列表
        this.$get(".shield").params({Addr: this.address + searchIp, NodeId: node}).success(resp => {
            if (resp.code === 200) {
                if (resp.data.shield)
                    this.shieldList = resp.data.shield
                else
                    this.shieldList = []
                this.tableState = state
            }
        })

    };
    this.onOpenConfig = function (addr) {
        this.address = addr
        let node = this.getNodeId()
        let that = this
        //ip直通 防护策略
        this.$get(".set").params({Addr: addr, NodeId: node}).success(resp => {
            if (resp.code === 200) {
                that.ignore = resp.data.ignore
                that.level = resp.data.level
                //屏蔽列表
                that.$get(".shield").params({Addr: that.address, NodeId: node}).success(resp => {
                    if (resp.code === 200) {
                        if (resp.data.shield)
                            that.shieldList = resp.data.shield
                        else
                            that.shieldList = []
                        that.changeState(2)
                    }
                })
            }
        })
    }

    //配置里面的列表切换
    this.changeListState = function (state) {
        if (this.tableState !== state) {
            if (state === 1) { //屏蔽列表
                this.shieldSearchList(state)
            } else {//连接列表
                this.linkSearchList(state)
            }
        }
    }

    //连接列表查询
    this.linkSearchList = function (state) {
        let searchIp = this.src_ip === '' ? '' : '-' + this.src_ip
        let node = this.getNodeId()
        //屏蔽列表
        this.$get(".link").params({Addr: this.address + searchIp, NodeId: node}).success(resp => {
            if (resp.code === 200) {
                this.linkList = resp.data.list
                this.tableState = state
            }
        })
    }
    //导出
    this.onExport = function () {
        var inputValue = ""
        if (this.tableState == 1) {
            inputValue = document.getElementById("linkHostInput").value
        } else {
            inputValue = document.getElementById("unlinkHostInput").value
        }
        console.log(inputValue)
        //todo
    }
    //全部释放 如果传入id 则单独释放 否则释放全部
    this.onRelease = function (item) {
        let adds = []
        let node = this.getNodeId()
        let msg = ''
        if (item === "all") {//全部释放
            adds[0] = this.address
            msg = "全部"
        } else {
            adds[0] = item.RemoteAddress + '-' + item.LocalAddress
        }
        //ip直通 防护策略
        teaweb.confirm("确定要" + msg + "释放吗？", function () {
            this.post(".shield").params({Addr: adds, NodeId: node}).success(resp => {
                if (resp.code === 200) {
                    teaweb.success("释放成功", function () {
                        teaweb.reload()
                    })
                }
            }).refresh()
        })
    }
})