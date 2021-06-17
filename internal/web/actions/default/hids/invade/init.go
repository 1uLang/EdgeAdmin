package invade

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//入侵威胁

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeHids)).
			Data("teaMenu", "hids").
			Prefix("/hids/risk").
			Get("", new(IndexAction)).
			//病毒木马
			GetPost("/virus", new(VirusAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/virus/detailList", new(VirusDetailListAction)). //病毒木马详情列表
			Get("/virus/detail", new(VirusDetailAction)).         //病毒木马详情
			//弱口令
			//GetPost("/weak", new(WeakAction)).
			//Get("/weakList", new(WeakListAction)).
			////风险账号
			//GetPost("/dangerAccount", new(DangerAccountAction)).
			//Get("/dangerAccountList", new(DangerAccountListAction)).
			////配置缺陷
			//GetPost("/configDefect", new(ConfigDefectAction)).
			//Get("/configDefectList", new(ConfigDefectAction)).
			EndAll()
	})
}
