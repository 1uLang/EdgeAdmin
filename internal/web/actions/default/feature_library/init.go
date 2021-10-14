package feature_library

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeConfig)).
			Helper(NewHelper("virus")).
			Data("teaMenu", "virus").
			Prefix("/feature_library/virus").
			Get("", new(IndexAction)).
			EndAll()
	})
}
