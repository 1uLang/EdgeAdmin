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
			Data("teaMenu", "scans").
			Prefix("/webscan/scans").
			Get("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			Get("/vulnerabilities", new(VulnerabilitiesAction)).
			Post("/delete", new(DeleteAction)).
			Post("/stop", new(StopAction)).
			Get("/statistics", new(StatisticsAction)).
			EndAll()
	})
}
