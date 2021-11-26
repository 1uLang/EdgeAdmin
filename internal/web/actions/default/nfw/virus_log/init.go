package virus_log

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/nfw/settingsutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNfw)).
			Helper(settingsutils.NewHelper("virusLog")).
			//Data("teaMenu", "virusLog").
			Prefix("/nfw/virusLog").
			Get("", new(VirusLogAction)).
			EndAll()
	})
}
