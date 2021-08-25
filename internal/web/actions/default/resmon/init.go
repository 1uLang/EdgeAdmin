package resmon

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "resmon").
			Prefix("/resmon").
			GetPost("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			Post("/delete", new(DeleteAction)).
			Post("/update", new(UpdateAction)).
			Get("/agents", new(AgentAction)).
			EndAll()
	})
}
