package configloaders

import (
	teaconst "github.com/TeaOSLab/EdgeAdmin/internal/const"
	"github.com/TeaOSLab/EdgeAdmin/internal/rpc"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/systemconfigs"
	"github.com/iwind/TeaGo/maps"
)

type AdminModuleCode = string

const (
	AdminModuleCodeDashboard AdminModuleCode = "dashboard" // 看板
	AdminModuleCodeServer    AdminModuleCode = "server"    // 网站
	AdminModuleCodeCerts     AdminModuleCode = "certs"     // 证书服务
	AdminModuleCodeNode      AdminModuleCode = "node"      // 节点
	AdminModuleCodeDNS       AdminModuleCode = "dns"       // DNS
	AdminModuleCodeNS        AdminModuleCode = "ns"        // 域名服务
	AdminModuleCodeAdmin     AdminModuleCode = "config"    // 系统用户
	AdminModuleCodeUser      AdminModuleCode = "user"      // 平台用户
	AdminModuleCodeFinance   AdminModuleCode = "finance"   // 财务
	AdminModuleCodePlan      AdminModuleCode = "plan"      // 套餐
	AdminModuleCodeLog       AdminModuleCode = "log"       // 日志
	AdminModuleCodeSetting   AdminModuleCode = "setting"   // 设置
	AdminModuleCodeAssembly  AdminModuleCode = "assembly"  // 只要登录就可以访问的模块
	AdminModuleCodeCommon    AdminModuleCode = "common"    // 只要登录就可以访问的模块
	AdminModuleCodeConfig    AdminModuleCode = "config"    // 平台管理
	AdminModuleCodeClusters  AdminModuleCode = "clusters"  // 边缘节点

	AdminModuleCodeDdos     AdminModuleCode = "ddos"      // ddos
	AdminModuleCodeWebScan  AdminModuleCode = "webscan"   // webscan
	AdminModuleCodeHids     AdminModuleCode = "hids"      // hids 主机防护
	AdminModuleCodeNfw      AdminModuleCode = "nfw"       // 下一代防火墙
	AdminModuleCodeMaltrail AdminModuleCode = "apt"       // apt检测
	AdminModuleCodeMonitor  AdminModuleCode = "monitor"   //监控告警
	AdminModuleCodeWAF      AdminModuleCode = "waf"       //web防火墙
	AdminModuleCodeAudit    AdminModuleCode = "audit"     //审计系统
	AdminModuleCodeFort     AdminModuleCode = "fortcloud" //堡垒机
	AdminModuleCodeBackup   AdminModuleCode = "backup"    //数据备份
	AdminModuleCodeHostlist AdminModuleCode = "hostlist"  //主机列表
)

var sharedAdminModuleMapping = map[int64]*AdminModuleList{} // adminId => AdminModuleList

func loadAdminModuleMapping() (map[int64]*AdminModuleList, error) {
	if len(sharedAdminModuleMapping) > 0 {
		return sharedAdminModuleMapping, nil
	}

	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		return nil, err
	}
	modulesResp, err := rpcClient.AdminRPC().FindAllAdminModules(rpcClient.Context(0), &pb.FindAllAdminModulesRequest{})
	if err != nil {
		return nil, err
	}
	mapping := map[int64]*AdminModuleList{}
	for _, m := range modulesResp.AdminModules {
		list := &AdminModuleList{
			IsSuper:  m.IsSuper,
			Fullname: m.Fullname,
			Theme:    m.Theme,
		}

		for _, pbModule := range m.Modules {
			list.Modules = append(list.Modules, &systemconfigs.AdminModule{
				Code:     pbModule.Code,
				AllowAll: pbModule.AllowAll,
				Actions:  pbModule.Actions,
			})
		}

		mapping[m.AdminId] = list
	}

	sharedAdminModuleMapping = mapping

	return sharedAdminModuleMapping, nil
}

func NotifyAdminModuleMappingChange() error {
	locker.Lock()
	defer locker.Unlock()
	sharedAdminModuleMapping = map[int64]*AdminModuleList{}
	_, err := loadAdminModuleMapping()
	return err
}

// CheckAdmin 检查用户是否存在
func CheckAdmin(adminId int64) bool {
	locker.Lock()
	defer locker.Unlock()

	// 如果还没有数据，则尝试加载
	if len(sharedAdminModuleMapping) == 0 {
		_, _ = loadAdminModuleMapping()
	}

	_, ok := sharedAdminModuleMapping[adminId]
	return ok
}

// AllowModule 检查模块是否允许访问
func AllowModule(adminId int64, module string) bool {
	locker.Lock()
	defer locker.Unlock()

	if module == AdminModuleCodeCommon {
		return true
	}

	if len(sharedAdminModuleMapping) == 0 {
		_, _ = loadAdminModuleMapping()
	}

	list, ok := sharedAdminModuleMapping[adminId]
	if ok {
		return list.Allow(module)
	}

	return false
}

// FindFirstAdminModule 获取管理员第一个可访问模块
func FindFirstAdminModule(adminId int64) (module AdminModuleCode, ok bool) {
	locker.Lock()
	defer locker.Unlock()
	list, ok2 := sharedAdminModuleMapping[adminId]
	if ok2 {
		if list.IsSuper {
			return AdminModuleCodeDashboard, true
		} else if len(list.Modules) > 0 {
			return list.Modules[0].Code, true
		}
	}
	return
}

// FindAdminFullname 查找某个管理员名称
func FindAdminFullname(adminId int64) string {
	locker.Lock()
	defer locker.Unlock()

	list, ok := sharedAdminModuleMapping[adminId]
	if ok {
		return list.Fullname
	}
	return ""
}

// FindAdminTheme 查找某个管理员选择的风格
func FindAdminTheme(adminId int64) string {
	locker.Lock()
	defer locker.Unlock()

	list, ok := sharedAdminModuleMapping[adminId]
	if ok {
		return list.Theme
	}
	return ""
}

// UpdateAdminTheme 设置某个管理员的风格
func UpdateAdminTheme(adminId int64, theme string) {
	locker.Lock()
	defer locker.Unlock()

	list, ok := sharedAdminModuleMapping[adminId]
	if ok {
		list.Theme = theme
	}
}

// AllModuleMaps 所有权限列表
func AllModuleMaps() []maps.Map {
	m := []maps.Map{
		{
			"name": "业务概览",
			"code": AdminModuleCodeDashboard,
			"url":  "/dashboard",
		},
		{
			"name": "平台管理",
			"code": AdminModuleCodeConfig,
			"url":  "/assembly",
		},
		{
			"name": "DDoS防护",
			"code": AdminModuleCodeDdos,
			"url":  "/ddos/host",
		},
		{
			"name": "云防火墙",
			"code": AdminModuleCodeNfw,
			"url":  "/nfw/nat",
		},
		{
			"name": "APT防御",
			"code": AdminModuleCodeMaltrail,
			"url":  "/apt/logs",
		},
	}
	if teaconst.IsPlus {
		m = append(m, maps.Map{
			"name": "智能DNS",
			"code": AdminModuleCodeNS,
			"url":  "/ns",
		})
	}
	m = append(m, []maps.Map{
		{
			"name": "WAF服务",
			"code": AdminModuleCodeServer,
			"url":  "/servers",
		},
		{
			"name": "证书服务",
			"code": AdminModuleCodeCerts,
			"url":  "/servers/certs",
		},
		{
			"name": "主机防护",
			"code": AdminModuleCodeHids,
			"url":  "/hids/examine",
		},
		{
			"name": "端点防护",
			"code": "nhids",
			"url":  "/hids/agents",
		},
		{
			"name": "漏洞扫描",
			"code": AdminModuleCodeWebScan,
			"url":  "/webscan/targets",
		},
		{
			"name": "云堡垒机",
			"code": AdminModuleCodeFort,
			"url":  "/fortcloud/assets",
		},
		{
			"name": "安全审计",
			"code": AdminModuleCodeAudit,
			"url":  "/audit/db",
		},
		{
			"name": "数据备份",
			"code": AdminModuleCodeBackup,
			"url":  "/databackup",
		},
		//{
		//	"name": "监控告警",
		//	"code": AdminModuleCodeMonitor,
		//	"url":  "/monitor",
		//},
		{
			"name": "系统设置",
			"code": AdminModuleCodeSetting,
			"url":  "/settings/server",
		},
	}...)
	return m
}
