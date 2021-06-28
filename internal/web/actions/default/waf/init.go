package waf

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.Prefix("/waf").
			Data("teaMenu", "waf").
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeWAF)).
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
