package assets

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeFort)).
			Data("teaMenu", "fort").
			Prefix("/fort/assets/system-users").
			Get("", new(IndexAction)).
			EndAll()
	})
}
