package file

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
			Prefix("/audit/form/file").
			Get("", new(IndexAction)).
			//GetPost("/createPopup", new(CreatePopupAction)).
			//GetPost("/ddos/host/shield_list", new(UpdatePopupAction)).
			EndAll()
	})
}
