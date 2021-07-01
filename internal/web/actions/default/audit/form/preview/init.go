package preview

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeAudit)).
			Data("teaMenu", "preview").
			Prefix("/audit/form/preview").
			Get("", new(IndexAction)).
			//GetPost("/createPopup", new(CreatePopupAction)).
			//GetPost("/ddos/host/shield_list", new(UpdatePopupAction)).
			EndAll()
	})
}
