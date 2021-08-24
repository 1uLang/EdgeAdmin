package helpers

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeAdmin/internal/const"
	"github.com/TeaOSLab/EdgeAdmin/internal/setup"
	"github.com/TeaOSLab/EdgeAdmin/internal/utils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"
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

	// 恢复模式
	if teaconst.IsRecoverMode {
		action.RedirectURL("/recover")
		return false
	}

	// DEMO模式
	if teaconst.IsDemoMode {
		if action.Request.Method == http.MethodPost {
			var actionName = action.Spec.ClassName[strings.LastIndex(action.Spec.ClassName, ".")+1:]
			var denyPrefixes = []string{"Update", "Create", "Delete", "Truncate", "Clean", "Clear", "Reset", "Add", "Remove", "Sync"}
			for _, prefix := range denyPrefixes {
				if strings.HasPrefix(actionName, prefix) {
					action.Fail(teaconst.ErrorDemoOperation)
					return false
				}
			}

			if strings.Index(action.Spec.PkgPath, "settings") > 0 || strings.Index(action.Spec.PkgPath, "delete") > 0 || strings.Index(action.Spec.PkgPath, "update") > 0 {
				action.Fail(teaconst.ErrorDemoOperation)
				return false
			}
		}
	}

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
	remoteAddr, _, _ := net.SplitHostPort(action.Request.RemoteAddr)
	if len(remoteAddr) > 0 && remoteAddr != action.RequestRemoteIP() && !checkIP(securityConfig, remoteAddr) {
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
	action.Data["teaTheme"] = configloaders.FindAdminTheme(adminId)

	action.Data["teaUserAvatar"] = ""

	if !action.Data.Has("teaMenu") {
		action.Data["teaMenu"] = ""
	}

	action.Data["teaModules"] = this.modules(adminId)

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
	action.Data["teaDemoEnabled"] = teaconst.IsDemoMode
	action.Data["teaShowFinance"] = configloaders.ShowFinance()
	if !action.Data.Has("teaSubMenu") {
		action.Data["teaSubMenu"] = ""
	}
	action.Data["teaCheckNodeTasks"] = configloaders.AllowModule(adminId, configloaders.AdminModuleCodeServer)
	action.Data["teaCheckDNSTasks"] = configloaders.AllowModule(adminId, configloaders.AdminModuleCodeServer)

	// 菜单
	action.Data["firstMenuItem"] = ""

	// 未读消息数
	action.Data["teaBadge"] = 0

	// 调用Init
	initMethod := reflect.ValueOf(actionPtr).MethodByName("Init")
	if initMethod.IsValid() {
		initMethod.Call([]reflect.Value{})
	}

	//fmt.Println("action.Request.RequestURI===",action.Request.RequestURI)
	//非检测请求  登录信息续期
	if !utils.UrlIn(action.Request.RequestURI) {
		res, _ := cache.GetInt(fmt.Sprintf("login_success_adminid_%v", adminId))
		if res == 0 {
			//30分钟没有操作  自动退出
			session.Delete()
			cache.DelKey(fmt.Sprintf("login_success_adminid_%v", adminId))
			this.login(action)
			return false
		}
		//续期
		cache.Incr(fmt.Sprintf("login_success_adminid_%v", adminId), time.Minute*30)
	}

	return true
}

// 菜单配置
func (this *userMustAuth) modules(adminId int64) []maps.Map {
	leftMaps := []maps.Map{
		{
			"code":   "dashboard",
			"module": configloaders.AdminModuleCodeDashboard,
			"name":   "业务概览",
			"icon":   "dashboard",
			"url":    "/dashboard",
		},
		{
			"code":   "assembly",
			"module": configloaders.AdminModuleCodeConfig,
			"name":   "平台管理",
			"icon":   "sitemap",
			"url":    "/assembly",
			"subItems": []maps.Map{
				{
					"name": "安全组件",
					"url":  "/assembly",
					"code": "assembly",
				},
				{
					"name": "资源监控",
					"url":  "/resmon",
					"code": "resmon",
				},
				{
					"name": "平台用户",
					"url":  "/users",
					"code": "users",
				},
				{
					"name": "系统用户",
					"url":  "/admins",
					"code": "admins",
				},
				{
					"name": "操作日志",
					"url":  "/log",
					"code": "log",
				},
				{
					"name": "通知媒介",
					"url":  "/admins/recipients",
					"code": "recipients",
				},
			},
		},
		{
			"code":   "ddos",
			"url":    "/ddos/host",
			"module": configloaders.AdminModuleCodeDdos,
			"name":   "DDoS防护",
			"icon":   "shield",
			"subItems": []maps.Map{
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
			"code":   "nfw",
			"url":    "/nfw/nat",
			"module": configloaders.AdminModuleCodeNfw,
			"name":   "云防火墙",
			"icon":   "bars",
			"subItems": []maps.Map{
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
				//{
				//	"name": "防病毒规则",
				//	"url":  "/nfw/virus",
				//	"code": "virus",
				//},
				{
					"name": "会话列表",
					"url":  "/nfw/conversation",
					"code": "conversation",
				},
				{
					"name": "安全警报",
					"url":  "/nfw/ips/alarm",
					"code": "alarm",
				},
			},
		},
		{
			"code":   "servers",
			"module": configloaders.AdminModuleCodeServer,
			"name":   "WAF服务",
			"icon":   "skyatlas",
			"url":    "/servers",
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
				//{
				//	"name": "日志策略",
				//	"url":  "/servers/accesslogs",
				//	"code": "accesslog",
				//	"isOn": teaconst.IsPlus,
				//},
				//{
				//	"name": "IP名单",
				//	"url":  "/servers/iplists",
				//	"code": "iplist",
				//},
				{
					"name": "集群列表",
					"url":  "/clusters",
					"code": "clusters",
				},
				{
					"name": "DNS服务商",
					"url":  "/dns/providers",
					"code": "providers",
				},
				//{
				//	"name": "统计指标",
				//	"url":  "/servers/metrics",
				//	"code": "metric",
				//},
			},
		},
		{
			"code": "certs",
			"name": "证书服务",
			"icon": "leanpub",
			"url":  "/servers/certs",
			"module": configloaders.AdminModuleCodeCerts,
		},
		{
			"code":   "hids",
			"url":    "/hids/examine",
			"module": configloaders.AdminModuleCodeHids,
			"name":   "主机防护",
			"icon":   "linux",
			"subItems": []maps.Map{
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
			"code":   "webscan",
			"url":    "/webscan/targets",
			"module": configloaders.AdminModuleCodeWebScan,
			"name":   "漏洞扫描",
			"icon":   "yelp",
			"subItems": []maps.Map{
				{
					"name": "扫描目标",
					"url":  "/webscan/targets",
					"code": "targets",
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
			"code":   "fortcloud",
			"url":    "/fortcloud/assets",
			"module": configloaders.AdminModuleCodeFort,
			"name":   "堡垒机",
			"icon":   "ioxhost",
			"subItems": []maps.Map{
				{
					"name": "资产管理",
					"url":  "/fortcloud/assets",
					"code": "assets",
					//"code":   "ns",
					//"module": configloaders.AdminModuleCodeNS,
					//"name":   "自建DNS",
					//"icon":   "cubes",
					//"isOn":   teaconst.IsPlus,
					//"subItems": []maps.Map{
					//	{
					//		"name": "域名管理",
					//		"url":  "/ns/domains",
					//		"code": "domain",
					//	},
					//	{
					//		"name": "集群管理",
					//		"url":  "/ns/clusters",
					//		"code": "cluster",
				},
				{
					"name": "授权凭证",
					"url":  "/fortcloud/cert",
					"code": "cert",
				},
				{
					"name": "会话管理",
					"url":  "/fortcloud/sessions",
					"code": "sessions",
				},
				{
					"name": "运维审计",
					"url":  "/fortcloud/audit",
					"code": "audit",
				},
				//{
				//	"name": "全局配置",
				//	"url":  "/ns/settings",
				//	"code": "setting",
				//},
				//{
				//	"name": "解析测试",
				//	"url":  "/ns/test",
				//	"code": "test",
				//},
			},
		},
		//{
		//	"code":   "clusters",
		//	"module": configloaders.AdminModuleCodeClusters,
		//	"name":   "边缘节点",
		//	"icon":   "instagram",
		//	"url":    "/clusters",
		//	"subItems": []maps.Map{
		//		{
		//			"name": "集群列表",
		//			"url":  "/clusters",
		//			"code": "clusters",
		//		},
		//		{
		//			"name": "节点日志",
		//			"url":  "/clusters/logs",
		//			"code": "logs",
		//		},
		//		{
		//			"name": "SSH认证",
		//			"url":  "/clusters/grants",
		//			"code": "grants",
		//		},
		//		{
		//			"name": "区域设置",
		//			"url":  "/clusters/regions",
		//			"code": "regions",
		//		},
		//	},
		//},
		//{
		//	"code":   "dns",
		//	"module": configloaders.AdminModuleCodeDNS,
		//	"name":   "域名解析",
		//	"icon":   "soundcloud",
		//	"url":    "/dns",
		//	"subItems": []maps.Map{
		//		{
		//			"name": "集群列表",
		//			"url":  "/dns",
		//			"code": "dns",
		//		},
		//		{
		//			"name": "问题修复",
		//			"url":  "/dns/issues",
		//			"code": "issues",
		//		},
		//		{
		//			"name": "DNS服务商",
		//			"url":  "/dns/providers",
		//			"code": "providers",
		//		},
		//	},
		//},
		//{
		//	"code":   "monitor",
		//	"module": configloaders.AdminModuleCodeMonitor,
		//	"name":   "监控告警",
		//	"icon":   "flickr ",
		//	"url":    "/monitor",
		//	"subItems": []maps.Map{
		//		{
		//			"name": "监控任务",
		//			"url":  "/monitor",
		//			"code": "monitor",
		//		},
		//		{
		//			"name": "告警通知",
		//			"url":  "/monitor/notice",
		//			"code": "notice",
		//		},
		//	},
		//},
		//{
		//	"code":   "audit",
		//	"module": configloaders.AdminModuleCodeAudit,
		//	"name":   "审计系统",
		//	"icon":   "",
		//	"subItems": []maps.Map{
		//		{
		//			"name": "全局状态",
		//			"url":  "/audit",
		//			"code": "",
		//		},
		//		{
		//			"name": "资产管理",
		//			"icon": "",
		//			"url":  "audit/assets/db",
		//			"subItems": []maps.Map{
		//				{
		//					"name": "数据库管理",
		//					"icon": "",
		//					"url":  "/audit/assets/db",
		//				},
		//				{
		//					"name": "主机管理",
		//					"icon": "",
		//					"url":  "/audit/assets/host",
		//				},
		//				{
		//					"name": "应用管理",
		//					"icon": "",
		//					"url":  "/audit/assets/app",
		//				},
		//				{
		//					"name": "agent管理",
		//					"icon": "",
		//					"url":  "/audit/assets/agent",
		{
			"code":   "audit",
			"url":    "/audit/db",
			"name":   "安全审计",
			"module": configloaders.AdminModuleCodeAudit,
			"icon":   "sellsy",
			"subItems": []maps.Map{
				{
					"name": "数据库管理",
					"url":  "/audit/db",
					"code": "db",
				},
				{
					"name": "主机管理",
					"url":  "/audit/host",
					"code": "host",
				},
				{
					"name": "应用管理",
					"url":  "/audit/app",
					"code": "app",
				},
				{
					"name": "审计日志",
					"url":  "/audit/logs",
					"code": "logs",
				},
				{
					"name": "订阅报告",
					"url":  "/audit/report",
					"code": "report",
				},
				{
					"name": "Agent管理",
					"url":  "/audit/agent",
					"code": "agent",
				},
			},
		}, {
			"code":   "databackup",
			"url":    "/databackup",
			"module": configloaders.AdminModuleCodeBackup,
			"name":   "数据备份",
			"icon":   "copy",
		},
		//		{
		//			"name": "查询分析",
		//			"icon": "",
		//			"url":  "/audit/query/logs",
		//			"subItems": []maps.Map{
		//				{
		//					"name": "审计日志",
		//					"icon": "",
		//					"url":  "/audit/query/logs",
		//				},
		//				{
		//					"name": "风险查询",
		//					"icon": "",
		//					"url":  "/audit/query/risk",
		//			"name": "授权管理",
		//			"code": "perms",
		//			"subItems": []maps.Map{
		//				{
		//					"name": "资产授权",
		//					"url":  "/fort/perms",
		//					"code": "perms",
		//				},
		//			},
		//		},
		//		{
		//			"name": "报表中心",
		//			"icon": "",
		//			"url":  "/audit/form/preview",
		//			"subItems": []maps.Map{
		//				{
		//					"name": "报表预览",
		//					"icon": "",
		//					"url":  "/audit/form/preview",
		//				},
		//				{
		//					"name": "报表订阅",
		//					"icon": "",
		//					"url":  "/audit/form/subscribe",
		//				},
		//				{
		//					"name": "审计归档",
		//					"icon": "",
		//					"url":  "/audit/form/file",
		//			"name": "会话管理",
		//			"code": "session",
		//			"subItems": []maps.Map{
		//				{
		//					"name": "会话管理",
		//					"url":  "/fort/session",
		//					"code": "session",
		//				},
		//				{
		//					"name": "命令记录",
		//					"url":  "/fort/session/command",
		//					"code": "command",
		//				},
		//			},
		//		},
		//	},
		//},
		//{
		//	"code":   "fortcloud",
		//	"module": configloaders.AdminModuleCodeFort,
		//	"name":   "堡垒机",
		//	"icon":   "history",
		//	"subItems": []maps.Map{
		//		{
		//			"name": "资产管理",
		//			"url":  "/fortcloud/assets",
		//			"code": "assets",
		//		},
		//		{
		//			"name": "管理账号",
		//			"url":  "/fortcloud/admins",
		//			"code": "admins",
		//		},
		//		{
		//			"name": "会话管理",
		//			"url":  "/fortcloud/sessions",
		//			"code": "sessions",
		//		},
		//		{
		//			"name": "运维审计",
		//			"url":  "/fortcloud/audit",
		//			"code": "audit",
		//		},
		//	},
		//},
		{
			"code":   "settings",
			"module": configloaders.AdminModuleCodeSetting,
			"name":   "系统设置",
			"icon":   "setting",
			"url":    "/settings/server",
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
	}
	leftResult := []maps.Map{}
	for _, m := range leftMaps {
		//去掉财务信息
		//if m.GetString("code") == "finance" && !configloaders.ShowFinance() {
		//	continue
		//}

		module := m.GetString("module")
		if configloaders.AllowModule(adminId, module) {
			leftResult = append(leftResult, m)
		}
	}
	return leftResult
}

// 跳转到登录页
func (this *userMustAuth) login(action *actions.ActionObject) {
	action.RedirectURL("/")
}
