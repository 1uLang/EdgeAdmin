package providers

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeDNS)).
			// 服务商
			Prefix("/dns/providers").
			Data("teaMenu", "providers").
			Data("teaSubMenu", "providers").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/updatePopup", new(UpdatePopupAction)).
			Post("/delete", new(DeleteAction)).
			Get("/provider", new(ProviderAction)).
			EndData().

			EndAll()
	})
}
