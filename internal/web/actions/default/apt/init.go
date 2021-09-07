package apt

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeMaltrail)).
			//Helper(settingsutils.NewHelper("virus")).
			Data("teaMenu", "apt").
			Prefix("/apt/logs").
			Get("", new(IndexAction)).
			EndAll()
	})
}
