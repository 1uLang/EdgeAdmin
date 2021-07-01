package db

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeAudit)).
			Data("teaMenu", "audit").
			Prefix("/audit/assets/db").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			EndAll()
	})
}
