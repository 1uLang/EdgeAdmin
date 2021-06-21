package logs

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNfw)).
			Data("teaMenu", "nfw").
			Prefix("/nfw/logs").
			Get("", new(IndexAction)).
			GetPost("/delete", new(DeleteAction)). //删除
			EndAll()
	})
}
