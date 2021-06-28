package helpers

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeAdmin/internal/const"
	"github.com/TeaOSLab/EdgeAdmin/internal/setup"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"net/http"
	"reflect"
)

// 认证拦截
type userMustAuth struct {
	AdminId int64
	module  string
}

func NewUserMustAuth(module string) *userMustAuth {
	return &userMustAuth{module: module}
}

func (this *userMustAuth) BeforeAction(actionPtr actions.ActionWrapper, paramName string) (goNext bool) {

	var action = actionPtr.Object()

	// 安全相关
	securityConfig, _ := configloaders.LoadSecurityConfig()
	if securityConfig == nil {
		action.AddHeader("X-Frame-Options", "SAMEORIGIN")
	} else if len(securityConfig.Frame) > 0 {
		action.AddHeader("X-Frame-Options", securityConfig.Frame)
	}
	action.AddHeader("Content-Security-Policy", "default-src 'self' data:; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'")

	// 检查IP
	if !checkIP(securityConfig, action.RequestRemoteIP()) {
		action.ResponseWriter.WriteHeader(http.StatusForbidden)
		return false
	}

	// 检查系统是否已经配置过
	if !setup.IsConfigured() {
		action.RedirectURL("/setup")
		return
	}

	var session = action.Session()
	var adminId = session.GetInt64("adminId")

	if adminId <= 0 {
		this.login(action)
		return false
	}

	// 检查用户是否存在
	if !configloaders.CheckAdmin(adminId) {
		session.Delete()

		this.login(action)
		return false
	}

	// 检查用户权限
	if len(this.module) > 0 && !configloaders.AllowModule(adminId, this.module) {
		action.ResponseWriter.WriteHeader(http.StatusForbidden)
		action.WriteString("Permission Denied.")
		return false
	}

	this.AdminId = adminId
	action.Context.Set("adminId", this.AdminId)

	if action.Request.Method != http.MethodGet {
		return true
	}

	config, err := configloaders.LoadAdminUIConfig()
	if err != nil {
		action.WriteString(err.Error())
		return false
	}

	// 初始化内置方法
	action.ViewFunc("teaTitle", func() string {
		return action.Data["teaTitle"].(string)
	})

	action.Data["teaShowVersion"] = config.ShowVersion
	action.Data["teaTitle"] = config.AdminSystemName
	action.Data["teaName"] = config.ProductName
	action.Data["teaFaviconFileId"] = config.FaviconFileId
	action.Data["teaLogoFileId"] = config.LogoFileId
	action.Data["teaUsername"] = configloaders.FindAdminFullname(adminId)

	action.Data["teaUserAvatar"] = ""

	if !action.Data.Has("teaMenu") {
		action.Data["teaMenu"] = ""
	}

	action.Data["teaModules"], action.Data["mainMenuData"] = this.modules(adminId)

	action.Data["teaSubMenus"] = []map[string]interface{}{}
	action.Data["teaTabbar"] = []map[string]interface{}{}
	if len(config.Version) == 0 {
		action.Data["teaVersion"] = teaconst.Version
	} else {
		action.Data["teaVersion"] = config.Version
	}
	action.Data["teaShowOpenSourceInfo"] = config.ShowOpenSourceInfo
	action.Data["teaIsSuper"] = false
	action.Data["teaIsPlus"] = teaconst.IsPlus
	action.Data["teaDemoEnabled"] = teaconst.IsDemo
	action.Data["teaShowFinance"] = configloaders.ShowFinance()
	if !action.Data.Has("teaSubMenu") {
		action.Data["teaSubMenu"] = ""
	}
	action.Data["teaCheckNodeTasks"] = configloaders.AllowModule(adminId, configloaders.AdminModuleCodeNode)
	action.Data["teaCheckDNSTasks"] = configloaders.AllowModule(adminId, configloaders.AdminModuleCodeDNS)

	// 菜单
	action.Data["firstMenuItem"] = ""

	// 未读消息数
	action.Data["teaBadge"] = 0

	// 调用Init
	initMethod := reflect.ValueOf(actionPtr).MethodByName("Init")
	if initMethod.IsValid() {
		initMethod.Call([]reflect.Value{})
	}

	return true
}

// 菜单配置
func (this *userMustAuth) modules(adminId int64) ([]maps.Map, []maps.Map) {
	leftMaps := []maps.Map{
		{
			"code":   "users",
			"module": configloaders.AdminModuleCodeUser,
			"name":   "平台用户",
			"icon":   "users",
			"subItems": []maps.Map{
				{
					"name": "平台用户",
					"url":  "/users",
					"code": "users",
				},
			},
		},
		{
			"code":     "admins",
			"module":   configloaders.AdminModuleCodeAdmin,
			"name":     "系统用户",
			"subtitle": "用户列表",
			"icon":     "user secret",
			"subItems": []maps.Map{
				{
					"name": "系统用户",
					"url":  "/admins",
					"code": "admins",
				},
				{
					"name": "通知媒介",
					"url":  "/admins/recipients",
					"code": "recipients",
					"isOn": teaconst.IsPlus,
				},
			},
		},
		{
			"code":   "log",
			"module": configloaders.AdminModuleCodeLog,
			"name":   "操作日志",
			"icon":   "history",
			"subItems": []maps.Map{
				{
					"name": "操作日志",
					"url":  "/log",
					"code": "log",
				},
			},
		},
		{
			"code":   "assembly",
			"module": configloaders.AdminModuleCodeAssembly,
			"name":   "节点管理",
			"icon":   "history",
			"subItems": []maps.Map{
				{
					"name": "节点管理",
					"url":  "/assembly",
					"code": "assembly",
				},
			},
		},
		{
			"code":   "settings",
			"module": configloaders.AdminModuleCodeSetting,
			"name":   "系统设置",
			"icon":   "setting",
			"subItems": []maps.Map{
				{
					"name": "基本设置",
					"url":  "/settings/server",
					"code": "server",
				},
				{
					"name": "高级设置",
					"url":  "/settings/advanced",
					"code": "advanced",
				},
			},
		},
		{
			"code":     "ddos",
			"module":   configloaders.AdminModuleCodeDdos,
			"name":     "DDOS",
			"subtitle": "全局状态",
			"icon":     "history",
			"subItems": []maps.Map{
				{
					"name": "全局状态",
					"url":  "/ddos",
					"code": "ddos",
				},
				{
					"name": "主机状态",
					"url":  "/ddos/host",
					"code": "host",
				},
				{
					"name": "黑白名单",
					"url":  "/ddos/whiteblacklist",
					"code": "whiteblacklist",
				},
				{
					"name": "统计日志",
					"url":  "/ddos/logs",
					"code": "logs",
				},
			},
		},
		{
			"code":     "webscan",
			"module":   configloaders.AdminModuleCodeWebScan,
			"name":     "漏洞扫描",
			"subtitle": "全局状态",
			"icon":     "setting",
			"subItems": []maps.Map{
				{
					"name": "全局状态",
					"url":  "/webscan",
					"code": "webscan",
				},
				{
					"name": "扫描目标",
					"url":  "/webscan/targets",
					"code": "targets",
				},
				{
					"name": "漏洞详情",
					"url":  "/webscan/vulnerabilities",
					"code": "vulnerabilities",
				},
				{
					"name": "扫描任务",
					"url":  "/webscan/scans",
					"code": "scans",
				},
				{
					"name": "扫描报告",
					"url":  "/webscan/reports",
					"code": "reports",
				},
			},
		},
		{
			"code":     "hids",
			"module":   configloaders.AdminModuleCodeHids,
			"name":     "主机安全防护",
			"subtitle": "全局状态",
			"icon":     "history",
			"subItems": []maps.Map{
				{
					"name": "全局状态",
					"url":  "/hids",
					"code": "hids",
				},
				{
					"name": "主机体检",
					"url":  "/hids/examine",
					"code": "examine",
				},
				{
					"name": "漏洞风险",
					"url":  "/hids/risk",
					"code": "risk",
				},
				{
					"name": "入侵威胁",
					"url":  "/hids/invade",
					"code": "invade",
				},
				{
					"name": "合规基线",
					"url":  "/hids/baseline",
					"code": "baseline",
				},
				{
					"name": "Agent管理",
					"url":  "/hids/agent",
					"code": "agent",
				},
			},
		},
		{
			"code":     "nfw",
			"module":   configloaders.AdminModuleCodeNfw,
			"name":     "nfw",
			"subtitle": "云防火墙",
			"icon":     "history",
			"subItems": []maps.Map{
				{
					"name": "全局状态",
					"url":  "/nfw",
					"code": "nfw",
				},
				{
					"name": "NAT规则",
					"url":  "/nfw/nat",
					"code": "nat",
				},
				{
					"name": "ACL规则",
					"url":  "/nfw/acl",
					"code": "acl",
				},
				{
					"name": "IPS规则",
					"url":  "/nfw/ips",
					"code": "ips",
				},
				{
					"name": "统计日志",
					"url":  "/nfw/logs",
					"code": "logs",
				},
			},
		},
		{
			"code":     "monitor",
			"module":   configloaders.AdminModuleCodeMonitor,
			"name":     "监控告警",
			"subtitle": "监控任务",
			"icon":     "history",
			"subItems": []maps.Map{
				{
					"name": "监控任务",
					"url":  "/monitor",
					"code": "monitor",
				},
				{
					"name": "告警通知",
					"url":  "/monitor/notice",
					"code": "notice",
				},
			},
		},
		{
			"code":   "waf",
			"module": configloaders.AdminModuleCodeWAF,
			"name":   "web防火墙",
			"icon":   "history",
			"subItems": []maps.Map{
				{
					"name": "全局状态",
					"url":  "/waf",
					"code": "waf",
				},
				{
					"name": "网站服务",
					"code": "servers",
					"icon": "history",
					"subItems": []maps.Map{
						{
							"name": "服务列表",
							"url":  "/servers",
							"code": "servers",
						},
						{
							"name": "通用设置",
							"url":  "/servers/components",
							"code": "components",
						},
						{
							"name": "服务分组",
							"url":  "/servers/components/groups",
							"code": "groups",
						},
						{
							"name": "缓存策略",
							"url":  "/servers/components/cache",
							"code": "cache",
						},
						{
							"name": "WAF策略",
							"url":  "/servers/components/waf",
							"code": "waf",
						},
						{
							"name": "证书管理",
							"url":  "/servers/certs",
							"code": "certs",
						},
					},
				},
				{
					"name": "边缘节点",
					"code": "clusters",
					"icon": "history",
					"subItems": []maps.Map{
						{
							"name": "集群列表",
							"url":  "/clusters",
							"code": "clusters",
						},
						{
							"name": "节点日志",
							"url":  "/clusters/logs",
							"code": "logs",
						},
						{
							"name": "SSH认证",
							"url":  "/clusters/grants",
							"code": "grants",
						},
						{
							"name": "区域设置",
							"url":  "/clusters/regions",
							"code": "regions",
						},
					},
				},
				{
					"name": "域名解析",
					"code": "dns",
					"icon": "history",
					"subItems": []maps.Map{
						{
							"name": "集群列表",
							"url":  "/dns",
							"code": "dns",
						},
						{
							"name": "问题修复",
							"url":  "/dns/issues",
							"code": "issues",
						},
						{
							"name": "DNS服务",
							"url":  "/dns/providers",
							"code": "providers",
						},
					},
				},
				{
					"name": "域名服务",
					"code": "ns",
					"icon": "history",
					"subItems": []maps.Map{
						{
							"name": "域名列表",
							"url":  "/ns",
							"code": "ns",
						},
						{
							"name": "集群管理",
							"url":  "/ns/clusters",
							"code": "clusters",
						},
						{
							"name": "线路管理",
							"url":  "/ns/routes",
							"code": "routes",
						},
						{
							"name": "访问日志",
							"url":  "/ns/clusters/accessLogs",
							"code": "accessLogs",
						},
						{
							"name": "运行日志",
							"url":  "/ns/clusters/logs",
							"code": "logs",
						},
					},
				},
				{
					"name": "通用设置",
					"url":  "/servers",
					"code": "servers",
				},
			},
		},
	}
	leftResult := []maps.Map{}
	for _, m := range leftMaps {
		if m.GetString("code") == "finance" && !configloaders.ShowFinance() {
			continue
		}

		module := m.GetString("module")
		if configloaders.AllowModule(adminId, module) {
			leftResult = append(leftResult, m)
		}
	}

	allMaps := []maps.Map{
		{
			"id":       1,
			"menuName": "首页",
			"pagePath": "/dashboard",
			"module":   configloaders.AdminModuleCodeDashboard,
			"leftMenu": []maps.Map{
				{
					"id":       1101,
					"leftName": "全局状态",
					"icon":     "",
					"pagePath": "/dashboard",
				},
			},
		},
		{
			"id":       2,
			"menuName": "节点管理",
			"pagePath": "/assembly",
			"module":   configloaders.AdminModuleCodeAssembly,
			"leftMenu": []maps.Map{
				{
					"id":       2101,
					"leftName": "节点管理",
					"icon":     "",
					"pagePath": "/assembly",
				},
			},
		},
		{
			"id":       3,
			"menuName": "安全服务",
			"module":   "safeServer",
			"dropItem": []maps.Map{
				{
					"id":       31,
					"dropName": "DDoS防火墙",
					"pagePath": "/ddos",
					"module":   configloaders.AdminModuleCodeDdos,
					"leftMenu": []maps.Map{
						{
							"id":       3101,
							"leftName": "全局状态",
							"icon":     "",
							"pagePath": "/ddos",
						},
						{
							"id":       3102,
							"leftName": "主机状态",
							"icon":     "",
							"pagePath": "/ddos/host",
						},
						{
							"id":       3103,
							"leftName": "黑白名单",
							"icon":     "",
							"pagePath": "/ddos/whiteblacklist",
						},
						{
							"id":       3104,
							"leftName": "统计日志",
							"icon":     "",
							"pagePath": "/ddos/logs",
						},
					},
				},
				{
					"id":       32,
					"dropName": "云防火墙",
					"pagePath": "/nfw",
					"module":   configloaders.AdminModuleCodeNfw,
					"leftMenu": []maps.Map{
						{
							"id":       3201,
							"leftName": "全局状态",
							"icon":     "",
							"pagePath": "/nfw",
						},
						{
							"id":       3202,
							"leftName": "NAT规则",
							"icon":     "",
							"pagePath": "/nfw/nat",
						},
						{
							"id":       3203,
							"leftName": "ACL规则",
							"icon":     "",
							"pagePath": "/nfw/acl",
						},
						{
							"id":       3204,
							"leftName": "IPS规则",
							"icon":     "",
							"pagePath": "/nfw/ips",
						},
						{
							"id":       3205,
							"leftName": "统计日志",
							"icon":     "",
							"pagePath": "/nfw/logs",
						},
					},
				},
				{
					"id":       33,
					"dropName": "WEB防火墙",
					"pagePath": "/waf",
					"module":   configloaders.AdminModuleCodeWAF,
					"leftMenu": []maps.Map{
						{
							"id":       3301,
							"leftName": "全局状态",
							"icon":     "",
							"pagePath": "/waf",
						},
						{
							"id":       3302,
							"leftName": "网站服务",
							"icon":     "",
							"pagePath": "/servers",
						},
						{
							"id":       3303,
							"leftName": "通用设置",
							"icon":     "",
							"pagePath": "/servers/components",
						},
						{
							"id":       3304,
							"leftName": "服务分组",
							"icon":     "",
							"pagePath": "/servers/components/groups",
						},
						{
							"id":       3305,
							"leftName": "缓存策略",
							"icon":     "",
							"pagePath": "/servers/components/cache",
						},
						{
							"id":       3306,
							"leftName": "WAF策略",
							"icon":     "",
							"pagePath": "/servers/components/waf",
						},
						{
							"id":       3307,
							"leftName": "证书管理",
							"icon":     "",
							"pagePath": "/servers/certs",
						},
						{
							"id":       3308,
							"leftName": "边缘节点",
							"icon":     "",
							"pagePath": "/clusters",
						},
						{
							"id":       3309,
							"leftName": "节点日志",
							"icon":     "",
							"pagePath": "/clusters/logs",
						},
						{
							"id":       3310,
							"leftName": "SSH认证",
							"icon":     "",
							"pagePath": "/clusters/grants",
						},
						{
							"id":       3311,
							"leftName": "区域设置",
							"icon":     "",
							"pagePath": "/clusters/regions",
						},
						{
							"id":       3312,
							"leftName": "域名解析",
							"icon":     "",
							"pagePath": "/dns",
						},
						{
							"id":       3313,
							"leftName": "问题修复",
							"icon":     "",
							"pagePath": "/dns/issues",
						},
						{
							"id":       3314,
							"leftName": "DNS服务",
							"icon":     "",
							"pagePath": "/clusters/providers",
						},
					},
				},
				{
					"id":       34,
					"dropName": "WEB漏洞扫描",
					"pagePath": "/webscan",
					"module":   configloaders.AdminModuleCodeWebScan,
					"leftMenu": []maps.Map{
						{
							"id":       3401,
							"leftName": "全局状态",
							"icon":     "",
							"pagePath": "/webscan",
						},
						{
							"id":       3402,
							"leftName": "扫描目标",
							"icon":     "",
							"pagePath": "/webscan/targets",
						},
						{
							"id":       3403,
							"leftName": "漏洞详情",
							"icon":     "",
							"pagePath": "/webscan/vulnerabilities",
						},
						{
							"id":       3404,
							"leftName": "扫描任务",
							"icon":     "",
							"pagePath": "/webscan/scans",
						},
						{
							"id":       3405,
							"leftName": "扫描报告",
							"icon":     "",
							"pagePath": "/webscan/reports",
						},
					},
				},
				{
					"id":       35,
					"dropName": "主机防护",
					"pagePath": "/hids",
					"module":   configloaders.AdminModuleCodeHids,
					"leftMenu": []maps.Map{
						{
							"id":       3501,
							"leftName": "全局状态",
							"icon":     "",
							"pagePath": "/hids",
						},
						{
							"id":       3502,
							"leftName": "主机体现",
							"icon":     "",
							"pagePath": "/hids/examine",
						},
						{
							"id":       3503,
							"leftName": "漏洞风险",
							"icon":     "",
							"pagePath": "/hids/risk",
						},
						{
							"id":       3504,
							"leftName": "入侵威胁",
							"icon":     "",
							"pagePath": "/hids/invade",
						},
						{
							"id":       3505,
							"leftName": "合规基线",
							"icon":     "",
							"pagePath": "/hids/baseline",
						},
						{
							"id":       3506,
							"leftName": "Agent管理",
							"icon":     "",
							"pagePath": "/hids/agent",
						},
					},
				},
			},
		},
		{
			"id":       4,
			"menuName": "监控告警",
			"pagePath": "/monitor",
			"module":   configloaders.AdminModuleCodeMonitor,
			"leftMenu": []maps.Map{
				{
					"id":       4101,
					"leftName": "监控任务",
					"icon":     "",
					"pagePath": "/monitor",
				},
				{
					"id":       4102,
					"leftName": "告警通知",
					"icon":     "",
					"pagePath": "/monitor/notice",
				},
				//{
				//	"id":       2101,
				//	"leftName": "告警设置",
				//	"icon":     "",
				//	"pagePath": "/assembly",
				//},
			},
		},
		{
			"id":       5,
			"menuName": "平台用户",
			"pagePath": "/users",
			"module":   configloaders.AdminModuleCodeUser,
			"leftMenu": []maps.Map{
				{
					"id":       2101,
					"leftName": "平台用户",
					"icon":     "",
					"pagePath": "/users",
				},
			},
		},
		{
			"id":       6,
			"menuName": "系统用户",
			"pagePath": "/admins",
			"module":   configloaders.AdminModuleCodeAdmin,
			"leftMenu": []maps.Map{
				{
					"id":       6101,
					"leftName": "系统用户",
					"icon":     "",
					"pagePath": "/admins",
				},
			},
		},
		{
			"id":       7,
			"menuName": "操作日志",
			"pagePath": "/log",
			"module":   configloaders.AdminModuleCodeLog,
			"leftMenu": []maps.Map{
				{
					"id":       7101,
					"leftName": "操作日志",
					"icon":     "",
					"pagePath": "/log",
				},
			},
		},
		{
			"id":       8,
			"menuName": "系统设置",
			"pagePath": "/settings/server",
			"module":   configloaders.AdminModuleCodeSetting,
			"leftMenu": []maps.Map{
				{
					"id":       8101,
					"leftName": "基本设置",
					"icon":     "",
					"pagePath": "/settings/server",
				},
				{
					"id":       8102,
					"leftName": "高级设置",
					"icon":     "",
					"pagePath": "/settings/database",
				},
			},
		},
	}
	result := []maps.Map{}
	for _, m := range allMaps {
		if m.GetString("code") == "finance" && !configloaders.ShowFinance() {
			continue
		}

		module := m.GetString("module")
		if module == "safeServer" { //安全服务
			safe := maps.Map{
				"id":       3,
				"menuName": "安全服务",
				"module":   "safeServer",
				"dropItem": []maps.Map{},
			}
			safeFlags := false
			for _, v := range m.GetSlice("dropItem") {
				md := v.(maps.Map)
				if configloaders.AllowModule(adminId, md.GetString("module")) {
					safeFlags = true
					safe["dropItem"] = append(safe.GetSlice("dropItem"), md)
				}
			}
			if safeFlags {
				result = append(result, safe)
			}
		} else if configloaders.AllowModule(adminId, module) {
			result = append(result, m)
		}
	}
	return leftResult, result
}

// 跳转到登录页
func (this *userMustAuth) login(action *actions.ActionObject) {
	action.RedirectURL("/")
}
