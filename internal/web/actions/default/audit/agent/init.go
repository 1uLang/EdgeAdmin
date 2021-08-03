package agent

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//审计-agent
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "agent").
			Prefix("/audit/agent").
			GetPost("", new(IndexAction)).
			//GetPost("/createPopup", new(CreatePopupAction)).
			//GetPost("/auth", new(AuthAction)).
			//GetPost("/delete", new(DeleteAction)).
			EndAll()
	})
}
