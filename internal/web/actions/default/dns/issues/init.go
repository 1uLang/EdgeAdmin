package issues

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			// 问题修复
			Prefix("/dns/issues").
			Data("teaMenu", "issues").
			Data("teaSubMenu", "issues").
			GetPost("", new(IndexAction)).
			GetPost("/updateNodePopup", new(UpdateNodePopupAction)).
			EndData().

			EndAll()
	})
}
