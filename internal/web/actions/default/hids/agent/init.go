package agent

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//agent管理

func init() {
	//_ = hids.InitAPIServer()
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "agent").
			Prefix("/hids/agent").
			GetPost("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			Get("/download", new(DownloadAction)).
			Get("/install", new(InstallAction)).
			Post("/disport", new(DisportAction)).
			Post("/delete", new(DeleteAction)).
			EndAll()
	})
}
