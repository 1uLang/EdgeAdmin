package host

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeDdos)).
			Data("teaMenu", "host").
			Prefix("/ddos/host").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)). //添加高防ip
			GetPost("/updatePopup", new(UpdatePopupAction)). //新增高防ip
			GetPost("/delete", new(DeleteAction)).           //删除高防ip
			GetPost("/shield", new(ShieldAction)).           //IP策略配置
			GetPost("/set", new(SetAction)).                 //IP策略配置
			Get("/link", new(LinkAction)).                   //IP策略配置
			EndAll()
	})
}
