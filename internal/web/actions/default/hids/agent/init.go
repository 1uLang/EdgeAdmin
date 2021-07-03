package agent

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//agent管理

func init() {
	_ = hids.InitAPIServer()
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeHids)).
			Data("teaMenu", "agent").
			Prefix("/hids/agent").
			Get("", new(IndexAction)).
			Get("/download", new(DownloadAction)).
			Get("/install", new(InstallAction)).
			Post("/disport", new(DisportAction)).
			EndAll()
	})
}
