package assembly

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//节点管理
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeConfig)).
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
	1: "DDoS防护",
	2: "云防火墙",
	3: "主机漏洞扫描",
	4: "WEB漏洞扫描",
	5: "主机防护",
	6: "安全审计",
	7: "堡垒机",
}
var idcMap = map[int]string{
	1: "成都IDC",
	2: "杭州IDC",
	3: "济南IDC",
}
