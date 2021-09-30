package channels

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeConfig)).
			Data("teaMenu", "channels").
			Prefix("/channels").
			GetPost("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			//GetPost("/update", new(UpdateAction)).
			//Post("/options", new(OptionsAction)).
			Post("/delete", new(DeleteAction)).
			EndAll()
	})
}
