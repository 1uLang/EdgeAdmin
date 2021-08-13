package databackup

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "databackup").
			Prefix("/databackup").
			GetPost("", new(IndexAction)).
			Get("/create", new(CreateAction)).
			Post("/delete", new(DeleteAction)).
			Get("/download", new(DownLoadAction)).
			EndAll()
	})
}