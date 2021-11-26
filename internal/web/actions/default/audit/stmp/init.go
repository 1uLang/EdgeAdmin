package stmp

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "stmp").
			Prefix("/audit/stmp").
			GetPost("", new(IndexAction)).
			Post("/check", new(CheckAction)).
			EndAll()
	})
}
