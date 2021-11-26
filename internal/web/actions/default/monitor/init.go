package monitor

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
			Prefix("/monitor").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)). //修改创建
			GetPost("/create", new(CreateAction)).           //修改创建
			GetPost("/delete", new(DeleteAction)).           //删除
			EndAll()
	})
}
