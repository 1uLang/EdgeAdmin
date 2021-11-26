package bwlist

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeHids)).
			Data("teaMenu", "bwlist").
			Prefix("/hids/bwlist").
			Get("", new(IndexAction)).
			GetPost("/del", new(DelAction)). //添删除ip
			GetPost("/createPopup", new(CreatePopupAction)).
			EndAll()
	})
}
