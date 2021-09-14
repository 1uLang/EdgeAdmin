Tea.context(function () {
    this.pageState = 1
    this.nSelectId = -1

    this.inputName = ""
    this.hostRule = ""
    this.mirrorRule = ""
    this.gridRule = ""
    this.diskRule = ""

    this.bShowDialog = false
    this.nShowRuleType = 1  //1 计算规格 2 镜像  3 网格   4 物理列表  5根云盘

    this.nPhysicsSelect = -1
    this.nMoreBtnSelet = -1

    this.nSelectHostRule = -1
    this.nSelectMirrorRule = -1
    this.nSelectDiskRule = -1
    this.sSelectGridRule = []
    this.nSelectGridRuleDataType = 1
    this.sShowGridData = []
    this.migrationUuid = -1   //需要迁移的云主机uuid
    this.specEditUuid = -1   //需要修改规格的云主机uuid

    let that = this
    this.$delay(function () {
        this.getSpec()
        this.getImage()
        this.getNetwork()
        this.getDisk()
    })

    //规格
    this.getSpec = async function () {
        await Tea.action("/hostlist/spec")
            .params({})
            .get()
            .success(function (resp) {
                let spec = resp.data.list
                this.hostRuleData = spec
            })
    }

    //镜像
    this.getImage = async function () {
        await Tea.action("/hostlist/image")
            .params({})
            .get()
            .success(function (resp) {
                let image = resp.data.list
                this.mirrorRuleData = image
            })
    }
    //根云盘 diskRuleData
    this.getDisk = async function () {
        await Tea.action("/hostlist/disk")
            .params({})
            .get()
            .success(function (resp) {
                let list = resp.data.list
                this.diskRuleData = list
            })
    }
    //网络
    this.getNetwork = async function () {
        await Tea.action("/hostlist/network")
            .params({})
            .get()
            .success(function (resp) {
                let network = resp.data.list
                this.gridRuleData = network
            })
    }


    //可迁移物理机列表
    this.getCandidate = async function (uuid) {
        await Tea.action("/hostlist/candidate")
            .params({
                uuid: uuid,
            })
            .get()
            .success(function (resp) {
                let list = resp.data.list
                this.physicsHost = list
            })
    }


    this.onChangeState = function (state) {
        if (this.pageState != state) {
            this.pageState = state
        }
    }
    this.getStatus = function (status) {
        switch (status) {
            case "Created":
                return "创建中"
            case "Starting":
                return "启动中"
            case "Running":
                return "运行中"
            case "Stopping":
                return "停止中"
            case "Stopped":
                return "已停止"
            case "Rebooting":
                return "重启中"
            case "Destroying":
                return "删除中"
            case "Destroyed":
                return "已删除"
            case "Migrating":
                return "迁移中"
            case "Expunging":
                return "清除中"
            case "Pausing":
                return "暂停中"
            case "Paused":
                return "已暂停"
            case "Resuming":
                return "恢复中"
            case "VolumeMigrating":
                return "批量迁移中"
            case "Unknown":
                return "未知"
            default:
                return "未知"
        }
    }

    this.onChangeTimeFormat = function (time) {
        // var resultTime = "";
        // if (time) {
        //     var tempTime = time.substring(0, time.indexOf("."));
        //     resultTime = tempTime.replace("T", " ");
        // }
        // return resultTime;
        var d = new Date(time);
        return d.toLocaleDateString() + " " + d.toLocaleTimeString()
    };

    this.onOpen = function (item) {
        if (item.status == 'Stopped') {
            let that = this
            that.$post(".activity")
                .params({
                    event: "start",
                    uuid: item.uuid,
                })
                .refresh()
        }
    }
    this.onStop = function (item) {

        if (item.status == 'Running' || item.status == 'Paused') {
            let that = this
            that.$post(".activity")
                .params({
                    event: "stop",
                    uuid: item.uuid,
                })
                .refresh()
        }
    }
    this.onMore = function (id) {
        this.nSelectId = id
    }

    this.onRestartPower = function (item) {
        if (item.status == 'Running') {
            this.onMore(-1)
            teaweb.confirm("确定要重启该主机？", function () {
                let that = this
                that.$post(".activity")
                    .params({
                        event: "restart",
                        uuid: item.uuid,
                    })
                    .refresh()
            })
        }
    }

    this.onPausePower = function (item) {
        if (item.status == 'Running') {
            this.onMore(-1)
            let that = this
            that.$post(".activity")
                .params({
                    event: "suspend",
                    uuid: item.uuid,
                })
                .refresh()
        }
    }

    this.onRebackPower = function (item) {
        console.log(item)
        if (item.status == 'Paused') {
            this.onMore(-1)
            let that = this
            that.$post(".activity")
                .params({
                    event: "unsuspend",
                    uuid: item.uuid,
                })
                .refresh()
        }
    }

    this.onMoveHost = function (item) {
        if (item.status != 'Stopping' && item.status != 'Rebooting' && item.status != 'Migrating' && !item.prohibitMigrating) {
            this.onMore(-1)
            this.onOpenPhysics(item.uuid)
            this.migrationUuid = item.uuid
        }
    }
    this.onEditHost = function (item) {
        if (item.status == 'Stopped') {
            this.onMore(-1)
            this.onOpenHostRule(item.uuid)
            this.specEditUuid = item.uuid
            // this.onSelectHostRule()
        }
    }
    this.onDeleteHost = function (item) {
        this.onMore(-1)
        teaweb.confirm("确定要删除该主机？", function () {
            let that = this
            that.$post(".activity")
                .params({
                    event: "delete",
                    uuid: item.uuid,
                })
                .refresh()
        })
    }

    this.onSave = function () {
        let that = this
        that.$post(".create")
            .params({
                Name: this.inputName,
                SpecUuid: this.nSelectHostRule,
                ImageUuid: this.nSelectMirrorRule,
                DiskUuid: this.nSelectDiskRule,
                NetworkUuid: this.sSelectGridRule,
            })
            .refresh()
    }

    this.onOpenPhysics = function (uuid) {
        this.nShowRuleType = 4
        this.nPhysicsSelect = -1
        this.bShowDialog = true
        this.getCandidate(uuid)
    }

    this.onOpenHostRule = function () {
        this.nShowRuleType = 1
        this.nSelectHostRule = -1
        this.bShowDialog = true

    }
    this.onOpenMirrorRule = function () {
        this.nShowRuleType = 2
        this.nSelectMirrorRule = -1
        this.bShowDialog = true
    }

    this.onOpenDiskRule = function () {
        this.nShowRuleType = 5
        this.nSelectDiskRule = -1
        this.bShowDialog = true
    }
    this.onOpenGridRule = function () {
        this.sSelectGridRule = []
        this.sShowGridData = this.gridRuleData.public
        this.nSelectGridRuleDataType = 1
        this.nShowRuleType = 3
        this.bShowDialog = true
    }
    this.onChangeGridRuleData = function (type) {
        if (this.nSelectGridRuleDataType != type) {
            this.nSelectGridRuleDataType = type
            this.sSelectGridRule = []
            if (type == 1) {
                this.sShowGridData = this.gridRuleData.public
            } else {
                this.sShowGridData = this.gridRuleData.private
            }

        }

    }

    this.onCloseRule = function () {
        this.bShowDialog = false
    }

    this.onGetItemInfo = function (id, table) {
        if (id && table && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].uuid == id) {
                    return table[index]
                }
            }
        }
        return null
    }
    //physics
    this.onSelectPhysics = function (id) {
        this.nPhysicsSelect = id
    }
    this.onSavePhysicsSelect = function () {
        var curSelectItem = this.onGetItemInfo(this.nPhysicsSelect, this.physicsHost)
        if (curSelectItem) {
            //req
            let that = this
            that.$post(".activity")
                .params({
                    event: "migration",
                    uuid: this.migrationUuid,
                    hostUuid: curSelectItem.uuid,
                })
                .refresh()
            this.onCloseRule()
        }
    }

    //host
    this.onSelectHostRule = function (id) {
        this.nSelectHostRule = id
    }
    this.onSelectDiskRule = function (id) {
        this.nSelectDiskRule = id
    }
    this.onSaveHostRuleSelect = function () {
        var curSelectItem = this.onGetItemInfo(this.nSelectHostRule, this.hostRuleData)
        if (curSelectItem) {
            this.hostRule = curSelectItem.name
            this.onCloseRule()

            //修改云主机的计算规格
            if (this.specEditUuid != -1) {
                let that = this
                that.$post(".activity")
                    .params({
                        event: "spec",
                        uuid: that.specEditUuid,
                        specUuid: curSelectItem.uuid,
                    })
                    .refresh()


                this.specEditUuid = -1
            }
        }


    }

    //mirror
    this.onSelectMirrorRule = function (id) {
        this.nSelectMirrorRule = id
    }
    this.onSaveMirrorRuleSelect = function () {
        var curSelectItem = this.onGetItemInfo(this.nSelectMirrorRule, this.mirrorRuleData)
        if (curSelectItem) {
            this.mirrorRule = curSelectItem.name
            this.onCloseRule()
        }
    }

    this.onSaveDiskRuleSelect = function () {
        var curSelectItem = this.onGetItemInfo(this.nSelectDiskRule, this.diskRuleData)
        if (curSelectItem) {
            this.diskRule = curSelectItem.name
            this.onCloseRule()
        }
    }

    //grid
    this.onRemoveTableItem = function (id) {
        if (id) {
            for (var index = 0; index < this.sSelectGridRule.length; index++) {
                if (this.sSelectGridRule[index] == id) {
                    this.sSelectGridRule.splice(index, 1)
                    break
                }
            }
        }
    }
    this.onCheckHadValue = function (id) {
        if (id) {
            for (var index = 0; index < that.sSelectGridRule.length; index++) {
                if (that.sSelectGridRule[index] == id) {
                    return true
                }
            }
        }
        return false
    }
    this.getShowImageName = function () {
        if (that.sSelectGridRule.length == 0) {
            return "/images/select_box.png"
        } else {
            let bAllSelect = true
            for (var index = 0; index < that.sShowGridData.length; index++) {
                if (!that.onCheckHadValue(that.sShowGridData[index].uuid)) {
                    bAllSelect = false
                    break
                }
            }
            if (bAllSelect) {
                return "/images/select_select.png"
            } else {
                return "/images/select_half_select.png"
            }
        }
    }
    this.getItemShowImageName = function (id) {
        if (that.onCheckHadValue(id)) {
            return "/images/select_select.png"
        } else {
            return "/images/select_box.png"
        }
    }
    this.onSelectGridRule = function (id) {
        if (this.onCheckHadValue(id)) {
            this.onRemoveTableItem(id)
        } else {
            this.sSelectGridRule.push(id)
        }
    }
    this.onSaveGridRuleSelect = function () {
        if (this.sSelectGridRule.length > 0) {
            this.gridRule = ""
            for (var index = 0; index < this.sSelectGridRule.length; index++) {
                var curSelectItem = this.onGetItemInfo(this.sSelectGridRule[index], this.sShowGridData)
                if (curSelectItem) {
                    this.gridRule = this.gridRule + "," + curSelectItem.name
                }
            }
            this.onCloseRule()
        }
    }
    this.onSelectAll = function () {
        console.log(11111)
        if (this.sSelectGridRule.length == 0) {
            this.sSelectGridRule = []
            for (var index = 0; index < this.sShowGridData.length; index++) {
                this.sSelectGridRule.push(this.sShowGridData[index].uuid)
            }
        } else {
            let bAllSelect = true
            for (var index = 0; index < that.sShowGridData.length; index++) {
                if (!this.onCheckHadValue(that.sShowGridData[index].uuid)) {
                    bAllSelect = false
                    break
                }
            }
            if (bAllSelect) {
                this.sSelectGridRule = []
            } else {
                this.sSelectGridRule = []
                for (var index = 0; index < this.sShowGridData.length; index++) {
                    this.sSelectGridRule.push(this.sShowGridData[index].uuid)
                }
            }
        }
    }
    this.mouseEnter = function (id) {
        this.nMoreBtnSelet = id
    }
    this.mouseLeave = function () {
        this.nMoreBtnSelet = -1
    }

    // this.tableData = [
    //     {
    //         uuid: 1,
    //         name: "云主机1",
    //         cpu: "4",
    //         memory: "4GB",
    //         ip: "192.168.0.1",
    //         physicsIP: "192.168.1.0",
    //         status: "Running",
    //         createTime: "2021-01-01T12:58:15.123"
    //     },
    //     {
    //         uuid: 2,
    //         name: "云主机1",
    //         cpu: "4",
    //         memory: "8GB",
    //         ip: "192.168.0.2",
    //         physicsIP: "192.168.2.0",
    //         status: 2,
    //         createTime: "2021-01-01T12:58:15.123"
    //     },
    //     {
    //         uuid: 3,
    //         name: "云主机1",
    //         cpu: "4",
    //         memory: "16GB",
    //         ip: "192.168.0.3",
    //         physicsIP: "192.168.3.0",
    //         status: 3,
    //         createTime: "2021-01-01T12:58:15.123"
    //     },
    //     {
    //         uuid: 4,
    //         name: "云主机1",
    //         cpu: "4",
    //         memory: "32GB",
    //         ip: "192.168.0.4",
    //         physicsIP: "192.168.4.0",
    //         status: 4,
    //         createTime: "2021-01-01T12:58:15.123"
    //     },
    // ]

    // this.hostRuleData = [
    //     {uuid: 1, name: "主机1", cpu: "4", memory: "4GB", createTime: "2021-01-01T12:58:15.123"},
    //     {uuid: 2, name: "主机2", cpu: "4", memory: "8GB", createTime: "2021-02-01T12:58:15.123"},
    //     {uuid: 3, name: "主机3", cpu: "4", memory: "16GB", createTime: "2021-03-01T12:58:15.123"},
    //     {uuid: 4, name: "主机4", cpu: "4", memory: "32GB", createTime: "2021-04-01T12:58:15.123"},
    // ]
    //
    // this.mirrorRuleData = [
    //     {
    //         uuid: 1,
    //         name: "CentOS7",
    //         mirrorType: "系统镜像1",
    //         mirrorFormat: "IOS",
    //         platform: "Linux",
    //         size: "950M",
    //         createTime: "2021-01-01T12:58:15.123"
    //     },
    //     {
    //         uuid: 2,
    //         name: "Ubuntu2.0",
    //         mirrorType: "系统镜像2",
    //         mirrorFormat: "WINDOW",
    //         platform: "Linux",
    //         size: "900M",
    //         createTime: "2021-02-01T12:58:15.123"
    //     },
    //     {
    //         uuid: 3,
    //         name: "CentOS8",
    //         mirrorType: "系统镜像3",
    //         mirrorFormat: "IOS1",
    //         platform: "Linux",
    //         size: "850M",
    //         createTime: "2021-03-01T12:58:15.123"
    //     },
    //     {
    //         uuid: 4,
    //         name: "CentOS9",
    //         mirrorType: "系统镜像4",
    //         mirrorFormat: "WINDOW2",
    //         platform: "Linux",
    //         size: "750M",
    //         createTime: "2021-04-01T12:58:15.123"
    //     },
    // ]
    //
    // this.gridRuleData = {
    //     public: [
    //         {uuid: 1, name: "L3Network-1", netType: "扁平网络", ip: "192.168.0.1/24", ipType: "IPv4", ipUse: "IPv4"},
    //         {uuid: 2, name: "L3Network-2", netType: "扁平网络", ip: "192.168.0.1/24", ipType: "IPv4", ipUse: "IPv4"},
    //         {uuid: 3, name: "L3Network-3", netType: "扁平网络", ip: "192.168.0.1/24", ipType: "IPv4", ipUse: "IPv4"},
    //         {uuid: 4, name: "L3Network-4", netType: "扁平网络", ip: "192.168.0.1/24", ipType: "IPv4", ipUse: "IPv4"},
    //     ],
    //     private: [
    //         {uuid: 1, name: "P3Network-1", netType: "扁平网络", ip: "192.168.0.1/24", ipType: "IPv4", ipUse: "IPv4"},
    //         {uuid: 2, name: "P3Network-2", netType: "扁平网络", ip: "192.168.0.1/24", ipType: "IPv4", ipUse: "IPv4"},
    //         {uuid: 3, name: "P3Network-3", netType: "扁平网络", ip: "192.168.0.1/24", ipType: "IPv4", ipUse: "IPv4"},
    //         {uuid: 4, name: "P3Network-4", netType: "扁平网络", ip: "192.168.0.1/24", ipType: "IPv4", ipUse: "IPv4"},
    //     ]
    //
    // }
    //
    this.physicsHost = [
        //     {uuid: 1, name: "Host-1", ip: "192.168.0.1/24", vrSkill: "KVM"},
        //     {uuid: 2, name: "Host-2", ip: "192.168.0.2/24", vrSkill: "KVM"},
        //     {uuid: 3, name: "Host-3", ip: "192.168.0.3/24", vrSkill: "KVM"},
        //     {uuid: 4, name: "Host-4", ip: "192.168.0.4/24", vrSkill: "KVM"},
    ]
    // this.diskRuleData = [
    //     {uuid: 1, name: "10G", size: "10G", createTime: "Sep 8, 2021 9:51:34 AM"},
    //     {uuid: 2, name: "40G", size: "40G", createTime: "Sep 8, 2021 9:51:34 AM"},
    // ]
})