Tea.context(function () {
    this.severity = 1

    this.nShowState = 1

    this.postId = 1     //接口* ID
    this.typeId = 1     //类型* ID
    this.outNewValue = ""//外部网络*的值
    this.sourceValue = ""//源的值
    this.sourceId = 1   //源的ID
    this.targetId = 1   //目标的ID
    this.descValue = "" //描述的ID

    this.onChangeShowState = function (state) {
        if (this.nShowState != state) {
            this.nShowState = state
            if (state == 2) {
                this.resetValue()
            }
        }
    }

    this.getStatus = function (status) {
        switch (status) {
            case '1':
                return "启已用"
            case '0':
                return "已停用"
            default:
                return "已停用"
        }
    }

    //修改配置
    this.onOpenChangeView = function (id) {
        console.log(id)
        for (var i = 0; i < this.tableData.length; i++) {
            if (this.tableData[i].id == id) {
                console.log("tableData")
                this.postId = this.tableData[i].postIndex
                this.typeId = this.tableData[i].typeIndex
                this.outNewValue = this.tableData[i].outIP
                this.sourceValue = this.tableData[i].sourceValue
                this.sourceId = this.tableData[i].sourceTypeIndex
                this.targetId = this.tableData[i].targetIndex
                this.descValue = this.tableData[i].desc
                break
            }
        }
        this.onChangeShowState(3)
    }

    //重置value
    this.resetValue = function () {
        this.postId = 1
        this.typeId = 1
        this.outNewValue = ""
        this.sourceValue = ""
        this.sourceId = 1
        this.targetId = 1
        this.descValue = ""
    }

    //保存配置
    this.onSaveConfig = function () {

    }

    //开启配置
    this.onOpenConfig = function (id) {

    }

    //关闭配置
    this.onOpenConfig = function (id,status) {
        let tops = "禁用"
        let statusUp = "0"
        let that = this
        if(status == "0"){
            tops = "启用"
            statusUp = "1"
        }
        teaweb.confirm("确定要"+tops+"此项吗？", function () {
            that.$post(".set")
                .params({
                    id: id,
                    status,statusUp,
                })
                .refresh()
        })
    }
    //删除配置
    this.onDelConfig = function (id) {
        teaweb.confirm("确定要删除所选规则？", function () {
            let that = this
            teaweb.confirm("确定要删除这个用户吗？", function () {
                that.$post(".delete")
                    .params({
                        userId: userId
                    })
                    .refresh()
            })
        })
    }

    this.dropNodeData = [
        {id: 1, value: 1},
        {id: 2, value: 2},
        {id: 3, value: 3},
        {id: 4, value: 4},
        {id: 5, value: 5},
        {id: 6, value: 6},
        {id: 7, value: 7},
        {id: 8, value: 8},
    ]

    this.hostData = [
        {id: 1, addr: "成都-ddos-192.168.1.1",},
        {id: 2, addr: "成都-ddos-192.168.1.2",},
        {id: 3, addr: "成都-ddos-192.168.1.3",},
        {id: 4, addr: "成都-ddos-192.168.1.4",},
    ]


    //postIndex:接口* 选择的id  typeIndex 类型* 选择的id  sourceTypeIndex 源* 选择的id  targetIndex 目标* 选择的id
    this.tableData = [
        {
            id: 1,
            postId: "LAN",
            intIP: "192.168.0.1",
            outIP: "192.168.1.1",
            targetIP: "192.168.3.1",
            desc: "描述",
            status: 1,
            postIndex: 1,
            typeIndex: 1,
            sourceTypeIndex: 1,
            targetIndex: 1,
            sourceValue: "源的值"
        },
        {
            id: 2,
            postId: "WAN",
            intIP: "192.168.0.2",
            outIP: "192.168.1.2",
            targetIP: "192.168.3.2",
            desc: "描述",
            status: 2,
            postIndex: 2,
            typeIndex: 1,
            sourceTypeIndex: 1,
            targetIndex: 1,
            sourceValue: "源的值"
        },
        {
            id: 3,
            postId: "LAN",
            intIP: "192.168.0.3",
            outIP: "192.168.1.3",
            targetIP: "192.168.3.3",
            desc: "描述",
            status: 1,
            postIndex: 1,
            typeIndex: 1,
            sourceTypeIndex: 1,
            targetIndex: 1,
            sourceValue: "源的值"
        },
        {
            id: 4,
            postId: "LAN",
            intIP: "192.168.0.4",
            outIP: "192.168.1.4",
            targetIP: "192.168.3.4",
            desc: "描述",
            status: 2,
            postIndex: 3,
            typeIndex: 1,
            sourceTypeIndex: 1,
            targetIndex: 1,
            sourceValue: "源的值"
        }
    ]

})