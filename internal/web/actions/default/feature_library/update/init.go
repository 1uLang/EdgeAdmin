package feature

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/feature_library"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeConfig)).
			Helper(feature_library.NewHelper("update")).
			Data("teaMenu", "virus").
			Prefix("/feature_library/update").
			Get("", new(IndexAction)).
			Prefix("/feature_library/auth_update").
			Get("", new(UpdateAction)).
			Prefix("/feature_library/status_update").
			Get("", new(UpdateStatusAction)).
			EndAll()
	})
}
