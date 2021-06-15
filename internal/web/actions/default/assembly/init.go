package assembly

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeAssembly)).
			Data("teaMenu", "assembly").
			Prefix("/assembly").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/update", new(UpdateAction)).
			Post("/idc/options", new(IDCOptionsAction)).
			Post("/options", new(OptionsAction)).
			Post("/delete", new(DeleteAction)).
			EndAll()
	})
}

var typeMap = map[int]string{
	0: "DDoS防护",
	1: "云防火墙",
	2: "主机防护",
	3: "WEB漏洞扫描",
	4: "主机漏洞扫描",
}
var idcMap = map[int]string{
	0: "成都IDC",
	1: "杭州IDC",
	2: "济南IDC",
}
