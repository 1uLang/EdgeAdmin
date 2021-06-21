package nat

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
			Prefix("/nfw/nat").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)). //修改创建
			GetPost("/delete", new(DeleteAction)).           //删除
			Post("/set", new(SetAction)).                    //禁用 启动
			EndAll()
	})
}
