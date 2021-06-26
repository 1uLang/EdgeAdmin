package reports

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
			Prefix("/webscan/reports").
			Get("", new(IndexAction)).
			Post("/create", new(CreateAction)).
			Post("/delete", new(DeleteAction)).
			Get("/download", new(DownloadAction)).
			EndAll()
	})
}