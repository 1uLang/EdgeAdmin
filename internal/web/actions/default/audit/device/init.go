package device

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//审计-设备
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "device").
			Prefix("/audit/device").
			GetPost("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/auth", new(AuthAction)).
			GetPost("/authorize", new(AuthorizeAction)).
			GetPost("/delete", new(DeleteAction)).
			EndAll()
	})
}
