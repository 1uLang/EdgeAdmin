package notice

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeMonitor)).
			Data("teaMenu", "monitor").
			Prefix("/monitor/notice").
			Get("", new(IndexAction)).
			GetPost("/delete", new(DeleteAction)). //删除
			EndAll()
	})
}
