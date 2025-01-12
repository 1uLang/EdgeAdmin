package sessions

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//web漏洞扫描
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "sessions").
			Prefix("/fortcloud/sessions").
			Get("", new(IndexAction)).
			Post("/disconnect", new(DisConnectAction)).
			//Post("/replay", new(ReplayAction)).
			//Post("/download", new(DownloadAction)).
			//Post("/monitor", new(MonitorAction)).
			EndAll()
	})
}
