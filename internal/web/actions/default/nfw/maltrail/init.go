package maltrail

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNfw)).
			//Helper(settingsutils.NewHelper("virus")).
			Data("teaMenu", "maltrail").
			Prefix("/nfw/maltrail/index").
			Get("", new(IndexAction)).
			EndAll()
	})
}
