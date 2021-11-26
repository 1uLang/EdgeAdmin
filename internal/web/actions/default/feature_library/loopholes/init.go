package loopholes

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
			Helper(feature_library.NewHelper("loopholes")).
			Data("teaMenu", "virus").
			Prefix("/feature_library/loopholes").
			Get("", new(IndexAction)).
			EndAll()
	})
}
