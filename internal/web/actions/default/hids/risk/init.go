package risk

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//漏洞风险

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeHids)).
			Data("teaMenu", "hids").
			Prefix("/hids/risk").
			Get("", new(IndexAction)).
			//系统漏洞
			GetPost("/systemRisk", new(SystemRiskAction)).
			Get("/systemRiskList", new(SystemRiskListAction)).
			//弱口令
			GetPost("/weak", new(WeakAction)).
			Get("/weakList", new(WeakListAction)).
			//风险账号
			GetPost("/dangerAccount", new(DangerAccountAction)).
			Get("/dangerAccountList", new(DangerAccountListAction)).
			//配置缺陷
			GetPost("/configDefect", new(ConfigDefectAction)).
			Get("/configDefectList", new(ConfigDefectAction)).
			EndAll()
	})
}
