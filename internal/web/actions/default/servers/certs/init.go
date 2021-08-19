package certs

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/certs/acme"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/certs/acme/users"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeCerts)).
			Helper(NewHelper()).
			Data("teaMenu", "certs").
			Data("teaSubMenu", "cert").
			Prefix("/servers/certs").
			Data("leftMenuItem", "cert").
			Get("", new(IndexAction)).
			GetPost("/uploadPopup", new(UploadPopupAction)).
			Post("/delete", new(DeleteAction)).
			GetPost("/updatePopup", new(UpdatePopupAction)).
			Get("/certPopup", new(CertPopupAction)).
			Get("/viewKey", new(ViewKeyAction)).
			Get("/viewCert", new(ViewCertAction)).
			Get("/downloadKey", new(DownloadKeyAction)).
			Get("/downloadCert", new(DownloadCertAction)).
			Get("/downloadZip", new(DownloadZipAction)).
			Get("/selectPopup", new(SelectPopupAction)).
			Get("/datajs", new(DatajsAction)).

			// ACME
			Prefix("/servers/certs/acme").
			Data("leftMenuItem", "acme").
			Get("", new(acme.IndexAction)).
			GetPost("/create", new(acme.CreateAction)).
			Post("/run", new(acme.RunAction)).
			GetPost("/updateTaskPopup", new(acme.UpdateTaskPopupAction)).
			Post("/deleteTask", new(acme.DeleteTaskAction)).
			Prefix("/servers/certs/acme/users").
			Get("", new(users.IndexAction)).
			GetPost("/createPopup", new(users.CreatePopupAction)).
			GetPost("/updatePopup", new(users.UpdatePopupAction)).
			Post("/delete", new(users.DeleteAction)).
			GetPost("/selectPopup", new(users.SelectPopupAction)).
			EndAll()
	})
}
