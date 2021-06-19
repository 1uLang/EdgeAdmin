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
			Data("teaMenu", "ddos").
			Prefix("/ddos/host").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)). //添加高防ip
			GetPost("/shieldList", new(ShieldListAction)).   //IP策略配置
			Post("/set", new(ShieldListAction)).             //IP策略配置
			EndAll()
	})
}