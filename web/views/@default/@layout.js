Tea.context(function () {
    this.moreOptionsVisible = false
    this.globalMessageBadge = 0
    this.showAPIAUTHVisible = -1

    if (typeof this.leftMenuItemIsDisabled == "undefined") {
        this.leftMenuItemIsDisabled = false
    }

    this.$delay(function () {
        if (this.$refs.focus != null) {
            this.$refs.focus.focus()
        }

        this.onSelectTopMenu(getSelectTopMenu(),getSelectTopDropId())

        // 检查消息
        this.checkMessages()

        // 检查集群节点同步
        this.loadNodeTasks();

        // 检查DNS同步
        this.loadDNSTasks()


    })

    /**
     * 左侧子菜单
     */
    this.showSubMenu = function (menu) {
        if (menu.alwaysActive) {
            return
        }
        if (this.teaSubMenus.menus != null && this.teaSubMenus.menus.length > 0) {
            this.teaSubMenus.menus.$each(function (k, v) {
                if (menu.id == v.id) {
                    return
                }
                v.isActive = false
            })
        }
        menu.isActive = !menu.isActive
    };

    /**
     * 检查消息
     */
    this.checkMessages = function () {
        this.$post("/messages/badge")
            .params({})
            .success(function (resp) {
                this.globalMessageBadge = resp.data.count
            })
            .done(function () {
                let delay = 6000
                if (this.globalMessageBadge > 0) {
                    delay = 30000
                }
                this.$delay(function () {
                    this.checkMessages()
                }, delay)
            })
    }

    this.checkMessagesOnce = function () {
        this.$post("/messages/badge")
            .params({})
            .success(function (resp) {
                this.globalMessageBadge = resp.data.count
            })
    }

    this.showMessages = function () {
        teaweb.popup("/messages", {
            height: "24em",
            width: "50em"
        })
    }

    /**
     * 底部伸展框
     */
    this.showQQGroupQrcode = function () {
        teaweb.popup("/about/qq", {
            width: "21em",
            height: "24em"
        })
    }

    /**
     * 弹窗中默认成功回调
     */
    if (window.IS_POPUP === true) {
        this.success = window.NotifyPopup
    }

    /**
     * 节点同步任务
     */
    this.doingNodeTasks = {
        isDoing: false,
        hasError: false,
        isUpdated: false
    }

    this.loadNodeTasks = function () {
        if (!Tea.Vue.teaCheckNodeTasks) {
            return
        }
        this.$post("/clusters/tasks/check")
            .success(function (resp) {
                this.doingNodeTasks.isDoing = resp.data.isDoing
                this.doingNodeTasks.hasError = resp.data.hasError
                this.doingNodeTasks.isUpdated = true
            })
            .done(function () {
                this.$delay(function () {
                    this.loadNodeTasks()
                }, 3000)
            })
    }

    this.showNodeTasks = function () {
        teaweb.popup("/clusters/tasks/listPopup", {
            height: "24em",
            width: "50em"
        })
    }

    /**
     * DNS同步任务
     */
    this.doingDNSTasks = {
        isDoing: false,
        hasError: false,
        isUpdated: false
    }

    this.loadDNSTasks = function () {
        if (!Tea.Vue.teaCheckDNSTasks) {
            return
        }
        this.$post("/dns/tasks/check")
            .success(function (resp) {
                this.doingDNSTasks.isDoing = resp.data.isDoing
                this.doingDNSTasks.hasError = resp.data.hasError
                this.doingDNSTasks.isUpdated = true
            })
            .done(function () {
                this.$delay(function () {
                    this.loadDNSTasks()
                }, 3000)
            })
    }

    this.showDNSTasks = function () {
        teaweb.popup("/dns/tasks/listPopup", {
            height: "24em",
            width: "50em"
        })
    }


    this.selectTopMenuId = 1    //选择顶部导航id
    this.selectLeftMenuId = 1   //选择左侧的菜单id
    this.letfMenuData = []


    this.getLeftData = function (id) {
        for(var i=0;i<this.mainMenuData.length;i++){
            if(this.mainMenuData[i].id==id){
                return this.mainMenuData[i].leftMenu
            }
        }
        return null
    }

    this.getDropLeftData = function (id,dropId) {
        for(var i=0;i<this.mainMenuData.length;i++){
            if(this.mainMenuData[i].id==id){
                if(this.mainMenuData[i].dropItem && this.mainMenuData[i].dropItem.length>0){
                    for(var j=0;j<this.mainMenuData[i].dropItem.length;j++){
                        if(this.mainMenuData[i].dropItem[j].id==dropId){
                            return this.mainMenuData[i].dropItem[j].leftMenu
                        }
                    }
                }
            }
        }
        return null
    }

    this.onSelectTopMenu=function (menuId,dropId) {
        console.log(menuId)
        if(this.selectTopMenuId !=menuId){
            this.selectTopMenuId = menuId
        }
        if(dropId && dropId>0){
            this.letfMenuData = this.getDropLeftData(menuId,dropId)
            
        }else{
            this.letfMenuData = this.getLeftData(menuId)
        }
        onSetTopMenuId(this.selectLeftMenuId)
        onSetSelectTopDropId(dropId)
        if(this.letfMenuData){
            this.onSelectLeftMenu(this.letfMenuData[0].id)
        }
        
    }
    this.onSelectLeftMenu=function (leftMenuId) { 
        if(this.selectLeftMenuId !=leftMenuId){
            this.selectLeftMenuId = leftMenuId
        }
        onSetLeftMenuId(this.selectLeftMenuId)
     }

    this.mainMenuData=[
        {
            id:1,
            menuName:"首页",
            pagePath:"/dashboard",
            leftMenu:[
                {id:1101,leftName:"全局状态",icon:"",pagePath:"/dashboard"}
            ]
        },
        {
            id:2,
            menuName:"节点管理",
            pagePath:"/assembly",
            leftMenu:[
                {id:2101,leftName:"节点管理",icon:"",pagePath:"/assembly"}
            ]
        },
        {
            id:3,
            menuName:"安全服务",
            dropItem:[
                {
                    id:31,
                    dropName:"DDos防火墙",
                    pagePath:"/ddos",
                    leftMenu:[
                        {id:3101,leftName:"全局状态",icon:"",pagePath:"/ddos"},
                        {id:3102,leftName:"主机状态",icon:"",pagePath:"/ddos/host"},
                        {id:3103,leftName:"黑白名单",icon:"",pagePath:"/ddos/whiteblacklist"},
                        {id:3104,leftName:"统计日志",icon:"",pagePath:"/ddos/logs"}
                    ]
                },
                {
                    id:32,
                    dropName:"云防火墙",
                    pagePath:"/nfw",
                    leftMenu:[
                        {id:3201,leftName:"全局状态",icon:"",pagePath:"/nfw"},
                        {id:3202,leftName:"NAT规则",icon:"",pagePath:"/nfw/nat"},
                        {id:3203,leftName:"ACL规则",icon:"",pagePath:"/nfw/acl"},
                        {id:3204,leftName:"IPS规则",icon:"",pagePath:"/nfw/ips"},
                        {id:3205,leftName:"统计日志",icon:"",pagePath:"/nfw/logs"}
                    ]
                },
                {
                    id:33,
                    dropName:"WEB防火墙",
                    pagePath:"/dashboard",
                    leftMenu:[
                        {id:3301,leftName:"全局状态",icon:"",pagePath:"/dashboard"},
                        {id:3302,leftName:"网站服务",icon:"",pagePath:"/servers"},
                        {id:3303,leftName:"通用设置",icon:"",pagePath:"/servers/components"},
                        {id:3304,leftName:"服务分组",icon:"",pagePath:"/servers/components/groups"},
                        {id:3305,leftName:"缓存策略",icon:"",pagePath:"/servers/components/cache"},
                        {id:3306,leftName:"WAF策略",icon:"",pagePath:"/servers/components/waf"},
                        {id:3307,leftName:"证书管理",icon:"",pagePath:"/servers/certs"},
                        {id:3308,leftName:"边缘节点",icon:"",pagePath:"/clusters"},
                        {id:3309,leftName:"节点日志",icon:"",pagePath:"/clusters/logs"},
                        {id:3310,leftName:"SSH认证",icon:"",pagePath:"/clusters/grants"},
                        {id:3311,leftName:"区域设置",icon:"",pagePath:"/clusters/regions"},
                        {id:3312,leftName:"域名解析",icon:"",pagePath:"/dns"},
                        {id:3313,leftName:"问题修复",icon:"",pagePath:"/dns/issues"},
                        {id:3314,leftName:"DNS服务",icon:"",pagePath:"/dns/providers"}
                    ]
                },
                {
                    id:34,
                    dropName:"WEB漏洞扫描",
                    pagePath:"/webscan",
                    leftMenu:[
                        {id:3401,leftName:"全局状态",icon:"",pagePath:"/webscan"},
                        {id:3402,leftName:"扫描目标",icon:"",pagePath:"/webscan/targets"},
                        {id:3403,leftName:"漏洞详情",icon:"",pagePath:"/webscan/vulnerabilities"},
                        {id:3404,leftName:"扫描任务",icon:"",pagePath:"/webscan/scans"},
                        {id:3405,leftName:"扫描报告",icon:"",pagePath:"/webscan/reports"}
                    ]
                },
                {
                    id:35,
                    dropName:"主机安全防护",
                    pagePath:"/hids",
                    leftMenu:[
                        {id:3501,leftName:"全局状态",icon:"",pagePath:"/hids"},
                        {id:3502,leftName:"主机体现",icon:"",pagePath:"/hids/examine"},
                        {id:3503,leftName:"漏洞风险",icon:"",pagePath:"/hids/risk"},
                        {id:3504,leftName:"入侵威胁",icon:"",pagePath:"/hids/invade"},
                        {id:3505,leftName:"合规基线",icon:"",pagePath:"/hids/baseline"},
                        {id:3506,leftName:"Agent管理",icon:"",pagePath:"/hids/agent"},
                    ]
                },
                // {
                //     id:36,
                //     dropName:"主机漏洞扫描",
                //     pagePath:"",
                //     leftMenu:[
                //         {id:3601,leftName:"平台首页",icon:"",pagePath:""}
                //     ]
                // },
                // {
                //     id:37,
                //     dropName:"安全审计",
                //     pagePath:"",
                //     leftMenu:[
                //         {id:3701,leftName:"全局状态",icon:"",pagePath:""},
                //         {id:3702,leftName:"资产管理",icon:"",pagePath:""},
                //         {id:3703,leftName:"数据库管理",icon:"",pagePath:""},
                //         {id:3704,leftName:"主机管理",icon:"",pagePath:""},
                //         {id:3705,leftName:"应用管理",icon:"",pagePath:""},
                //         {id:3706,leftName:"agent管理",icon:"",pagePath:""},
                //         {id:3707,leftName:"查询分析",icon:"",pagePath:""},
                //         {id:3708,leftName:"审计日志",icon:"",pagePath:""},
                //         {id:3709,leftName:"风险查询",icon:"",pagePath:""},
                //         {id:3710,leftName:"报表中心",icon:"",pagePath:""},
                //         {id:3711,leftName:"审计归档",icon:"",pagePath:""},
                //     ]
                // },
                // {
                //     id:38,
                //     dropName:"堡垒机",
                //     pagePath:"",
                //     leftMenu:[
                //         {id:3801,leftName:"平台首页",icon:"",pagePath:""}
                //     ]
                // }
            ],
        },
        {
            id:4,menuName:"监控告警",
            pagePath:"/monitor",
            leftMenu:[
                {id:4101,leftName:"监控任务",icon:"",pagePath:"/monitor"},
                {id:4102,leftName:"告警通知",icon:"",pagePath:"/monitor/notice"},
                // {id:4103,leftName:"告警设置",icon:"",pagePath:""},
            ]
        },
        {
            id:5,
            menuName:"平台用户",
            pagePath:"/users",
            leftMenu:[
                {id:5101,leftName:"平台用户",icon:"",pagePath:"/users"}
            ]
        },
        {
            id:6,
            menuName:"系统用户",
            pagePath:"/admins",
            leftMenu:[
                {id:6101,leftName:"系统用户",icon:"",pagePath:"/admins"}
            ]
        },
        {
            id:7,
            menuName:"操作日志",
            pagePath:"/log",
            leftMenu:[
                {id:7101,leftName:"操作日志",icon:"",pagePath:"/log"}
            ]
        },
        {
            id:8,
            menuName:"系统设置",
            pagePath:"/settings/server",
            leftMenu:[
                {id:8101,leftName:"基本设置",icon:"",pagePath:"/settings/server"},
                {id:8102,leftName:"高级设置",icon:"",pagePath:"/settings/database"}
            ]
        },
    ]
});

window.NotifySuccess = function (message, url, params) {
    if (typeof (url) == "string" && url.length > 0) {
        if (url[0] != "/") {
            url = Tea.url(url, params);
        }
    }
    return function () {
        teaweb.success(message, function () {
            window.location = url;
        });
    };
};

window.NotifyReloadSuccess = function (message) {
    return function () {
        teaweb.success(message, function () {
            window.location.reload()
        })
    }
}

window.NotifyDelete = function (message, url, params) {
    teaweb.confirm(message, function () {
        Tea.Vue.$post(url)
            .params(params)
            .refresh();
    });
};

window.NotifyPopup = function (resp) {
    window.parent.teaweb.popupFinish(resp);
};

window.ChangePageSize = function (size) {
    let url = window.location.toString();
    if (url.indexOf("pageSize") > 0) {
        url = url.replace(/pageSize=\d+/g, "pageSize=" + size);
    } else {
        if (url.indexOf("?") > 0) {
            url += "&pageSize=" + size;
        } else {
            url += "?pageSize=" + size;
        }
    }
    window.location = url;
};