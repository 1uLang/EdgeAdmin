package logs

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeDdos)).
			Data("teaMenu", "logs").
			Prefix("/ddos/logs").
			Get("", new(IndexAction)).
			Get("/attacks", new(AttacksAction)).
			Get("/link", new(LinkAction)).
			EndAll()
	})
}
