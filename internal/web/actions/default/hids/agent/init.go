package agent

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//agent管理

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeHids)).
			Data("teaMenu", "hids").
			Prefix("/hids/agent").
			Get("", new(IndexAction)).
			Get("/download", new(DisportAction)).
			Get("/install", new(DisportAction)).
			Post("/disport", new(DisportAction)).
			EndAll()
	})
}
