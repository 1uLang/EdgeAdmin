package examine

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//主机体检
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeHids)).
			Data("teaMenu", "examine").
			Prefix("/hids/examine").
			Get("", new(IndexAction)).
			Get("/detail", new(DetailAction)).
			Post("/scans", new(ScanAction)).
			EndAll()
	})
}
