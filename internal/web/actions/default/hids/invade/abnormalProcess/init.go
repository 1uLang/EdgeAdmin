package abnormalProcess

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
			Prefix("/hids/invade/abnormalProcess").
			GetPost("", new(IndexAction)).
			Get("/detailList", new(DetailListAction)). //异常进程 详情列表
			Get("/detail", new(DetailAction)).         //异常进程 详情

			EndAll()
	})
}
