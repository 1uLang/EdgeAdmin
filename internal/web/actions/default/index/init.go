package index

import (
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Post("/checkOTP", new(CheckOTPAction)).
			Prefix("/").
			GetPost("", new(IndexAction)).
			GetPost("updatePwd", new(UpdatePwdAction)).
			Post("renewal", new(RenewalAction)).
			EndAll()
	})
}
