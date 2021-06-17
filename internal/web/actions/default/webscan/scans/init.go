package scans

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//漏洞

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeWebScan)).
			Data("teaMenu", "webscan").
			Prefix("/webscan/scans").
			Get("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			Post("/delete", new(DeleteAction)).
			Post("/stop", new(StopAction)).
			EndAll()
	})
}
