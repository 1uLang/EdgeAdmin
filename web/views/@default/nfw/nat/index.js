Tea.context(function () {
    this.severity = 1

    this.nShowState = 1

    this.id = "" //id
    this.interface = "wan"     //接口* ID
    this.interfaces = []     //接口* 下拉列表
    this.type = "binat"     //类型* ID
    this.types = []     //类型* 下拉列表
    this.external = ""//外部网络*的值
    this.src = ""//源的值
    this.srcmask = ""//源的掩码
    this.dst = ""   //目标
    this.dstinput = ""   //目标 input
    this.dsts = []   //目标 下拉select
    this.dstmask = ""   //目标的掩码
    this.descr = "" //描述
    // this.selectNode = 0 //节点


    this.$delay(function () {
        let curSelectNode = localStorage.getItem("nfwSelectNodeId");
        if(curSelectNode){
            this.selectNode = curSelectNode
        }
    })

    this.onChangeShowState = function (state) {
        if (this.nShowState != state) {
            this.nShowState = state
            if (state == 2) {
                // this.resetValue()
            }
        }
    }

    this.getStatus = function (status) {
        switch (status) {
            case '1':
                return "已启用"
            case '0':
                return "已停用"
            default:
                return "已停用"
        }
    }

    //修改配置
    this.onOpenChangeView = function (id) {
        //获取详细数据
        this.GetNatInfo(id)
    }

    //保存配置
    this.onSaveConfig = function () {

        Tea.action("/nfw/nat/createPopup")
            .params({
                nodeId: this.selectNode,
                id: this.id,
                interface: this.interface,
                type: this.type,
                external: this.external,
                src: this.src,
                srcmask: this.srcmask,
                // dst: this.dst,
                dstinput: this.dstinput,
                dstmask: this.dstmask,
                descr: this.descr,
            })
            .post()
            .success(function (resp) {
            }).refresh()
    }


    //关闭配置
    this.onOpenConfig = function (id, status) {
        let tops = "禁用"
        let statusUp = "0"
        let that = this
        if (status == "0") {
            tops = "启用"
            statusUp = "1"
        }
        teaweb.confirm("确定要" + tops + "此项吗？", function () {
            that.$post(".set")
                .params({
                    id: id,
                    status: statusUp,
                    nodeId: this.selectNode,
                })
                .refresh()
        })
    }
    //删除配置
    this.onDelConfig = function (id) {
        teaweb.confirm("确定要删除所选规则？", function () {
            let that = this
            that.$post(".delete")
                .params({
                    id: id,
                    nodeId: this.selectNode,
                })
                .refresh()
        })
    }

    this.masks = [
        {id: 1, value: 1}, {id: 2, value: 2}, {id: 3, value: 3}, {id: 4, value: 4},
        {id: 5, value: 5}, {id: 6, value: 6}, {id: 7, value: 7}, {id: 8, value: 8},
        {id: 9, value: 9}, {id: 10, value: 10}, {id: 11, value: 11}, {id: 12, value: 12},
        {id: 13, value: 13}, {id: 14, value: 14}, {id: 15, value: 15}, {id: 16, value: 16},
        {id: 17, value: 17}, {id: 18, value: 18}, {id: 19, value: 19}, {id: 20, value: 20},
        {id: 21, value: 21}, {id: 22, value: 22}, {id: 23, value: 23}, {id: 24, value: 24},
        {id: 25, value: 25}, {id: 26, value: 26}, {id: 27, value: 27}, {id: 28, value: 28},
        {id: 29, value: 29}, {id: 30, value: 30}, {id: 31, value: 31}, {id: 32, value: 32},
    ]

    //获取当前选中的节点
    this.GetSelectNode = function (event) {
        this.selectNode = event.target.value; //获取option对应的value值
        localStorage.setItem("nfwSelectNodeId", this.selectNode);
        let node = this.selectNode
        window.location.href = '/nfw/nat?nodeId=' + node

    }

    //通过ID 获取详细数据
    this.GetNatInfo = async function (id) {

        await Tea.action("/nfw/nat/detail/options")
            .params({
                NodeId: this.selectNode,
                id: id,
                "act": "get-info"
            })
            .get()
            .success(function (resp) {

                this.srcmask = "32"
                let srcmask = resp.data.srcmask
                for (var i = 0; i < srcmask.length; i++) {
                    if (srcmask[i].selected == true) {
                        this.srcmask = srcmask[i].value

                    }
                }

                this.dstmask = "32"
                let dstmask = resp.data.dstmask
                for (var i = 0; i < dstmask.length; i++) {
                    if (dstmask[i].selected == true) {
                        this.dstmask = dstmask[i].value
                    }
                }
                this.id = id
                if (resp.data.interface.length > 0) {
                    this.interfaces = resp.data.interface

                    for (var i = 0; i < resp.data.interface.length; i++) { //接口下拉
                        if (resp.data.interface[i].selected == true) {
                            this.interface = resp.data.interface[i].value
                        }
                    }
                }

                if (resp.data.type.length > 0) {
                    this.types = resp.data.type
                    for (var i = 0; i < resp.data.type.length; i++) { //类型下拉
                        if (resp.data.type[i].selected == true) {
                            this.type = resp.data.type[i].value
                        }
                    }
                }
                this.external = resp.data.external
                if (id !== "")
                    this.src = resp.data.src
                else
                    this.src = ""
                // this.srcmask = this.tableDataList[i].srcmask //
                if (resp.data.dst.length > 0) {
                    this.dsts = resp.data.dst
                    for (var i = 0; i < resp.data.dst.length; i++) { //目标下拉
                        if (resp.data.dst[i].selected == true) {
                            this.dst = resp.data.dst[i].value
                            this.dstinput = resp.data.dst[i].value
                            if (resp.data.dst[i].data_other == true && !this.checkDstSrcType(this.dstinput)) {
                                //单个主机或网络
                                resp.data.dst[i].value = ""
                            }
                        }
                    }
                }
                // this.dstmask = this.tableDataList[i].dstmask //
                this.descr = resp.data.descr
                if (id != "") {
                    this.onChangeShowState(3)
                } else {
                    this.onChangeShowState(2)
                }


            })


    }


    //判断输入是否是输入类型
    this.checkDstSrcType = function (value) {
        let vmap = ["bogons", "bogonsv6", "virusprot", "sshlockout", "any", "(self)", "lan", "lanip",
            "lo0", "wan", "wanip"]
        if (vmap.indexOf(value) == -1) {
            //输入类型
            return true
        }
        return false
    }
})