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
	AdminModuleCodeNode      AdminModuleCode = "node"      // 节点
	AdminModuleCodeDNS       AdminModuleCode = "dns"       // DNS
	AdminModuleCodeNS        AdminModuleCode = "ns"        // 域名服务
	AdminModuleCodeAdmin     AdminModuleCode = "admin"     // 系统用户
	AdminModuleCodeUser      AdminModuleCode = "user"      // 平台用户
	AdminModuleCodeFinance   AdminModuleCode = "finance"   // 财务
	AdminModuleCodeLog       AdminModuleCode = "log"       // 日志
	AdminModuleCodeSetting   AdminModuleCode = "setting"   // 设置
	AdminModuleCodeAssembly  AdminModuleCode = "assembly"  // 只要登录就可以访问的模块
	AdminModuleCodeCommon    AdminModuleCode = "common"    // 只要登录就可以访问的模块
	AdminModuleCodeDdos      AdminModuleCode = "ddos"      // ddos
	AdminModuleCodeWebScan   AdminModuleCode = "webscan"   // webscan
	AdminModuleCodeHids      AdminModuleCode = "hids"      // hids 主机防护
	AdminModuleCodeNfw       AdminModuleCode = "nfw"       // 下一代防火墙
	AdminModuleCodeMonitor   AdminModuleCode = "monitor"   //监控告警
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

// AllModuleMaps 所有权限列表
func AllModuleMaps() []maps.Map {
	m := []maps.Map{
		{
			"name": "看板",
			"code": AdminModuleCodeDashboard,
			"url":  "/dashboard",
		},
		{
			"name": "网站服务",
			"code": AdminModuleCodeServer,
			"url":  "/servers",
		},
		{
			"name": "边缘节点",
			"code": AdminModuleCodeNode,
			"url":  "/clusters",
		},
		{
			"name": "域名解析",
			"code": AdminModuleCodeDNS,
			"url":  "/dns",
		},
	}
	if teaconst.IsPlus {
		m = append(m, maps.Map{
			"name": "域名服务",
			"code": AdminModuleCodeNS,
			"url":  "/ns",
		})
	}
	m = append(m, []maps.Map{
		{
			"name": "平台用户",
			"code": AdminModuleCodeUser,
			"url":  "/users",
		},
		{
			"name": "系统用户",
			"code": AdminModuleCodeAdmin,
			"url":  "/admins",
		},
		{
			"name": "财务管理",
			"code": AdminModuleCodeFinance,
			"url":  "/finance",
		},
		{
			"name": "日志审计",
			"code": AdminModuleCodeLog,
			"url":  "/log",
		},
		{
			"name": "系统设置",
			"code": AdminModuleCodeSetting,
			"url":  "/settings",
		},
		{
			"name": "DDoS防火墙",
			"code": AdminModuleCodeDdos,
			"url":  "/ddos",
		},
		{
			"name": "云防火墙",
			"code": AdminModuleCodeNfw,
			"url":  "/ddos",
		},
		{
			"name": "WEB防火墙",
			"code": AdminModuleCodeDashboard,
			"url":  "/dashboard",
		},
		{
			"name": "WEB漏洞扫描",
			"code": AdminModuleCodeWebScan,
			"url":  "/webscan",
		},
		{
			"name": "主机防护",
			"code": AdminModuleCodeHids,
			"url":  "/hids",
		},
	}...)
	return m
}
