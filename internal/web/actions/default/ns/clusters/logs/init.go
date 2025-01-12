package logs

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeWAF)).
			Data("teaMenu", "waf").
			Data("teaSubMenu", "log").
			Prefix("/ns/clusters/logs").
			Get("", new(IndexAction)).
			EndAll()
	})
}
