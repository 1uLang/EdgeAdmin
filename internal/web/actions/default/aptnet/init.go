package aptnet

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/apt_helper"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeMaltrail)).
			Helper(apt_helper.NewHelper("net")).
			Data("teaMenu", "apt").
			Prefix("/apt").
			Get("/net", new(NetAction)).
			EndAll()
	})
}
