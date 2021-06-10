package assembly

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeAssembly)).
			Data("teaMenu", "assembly").
			Prefix("/assembly").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			Get("/user", new(UserAction)).
			GetPost("/update", new(UpdateAction)).
			Post("/idc/options", new(DeleteAction)).
			GetPost("/features", new(FeaturesAction)).
			EndAll()
	})
}
