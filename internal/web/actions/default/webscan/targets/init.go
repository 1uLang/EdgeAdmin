package targets

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeWebScan)).
			Data("teaMenu", "webscan").
			Prefix("/webscan/targets").
			Get("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			Post("/delete", new(DeleteAction)).
			EndAll()
	})
}
