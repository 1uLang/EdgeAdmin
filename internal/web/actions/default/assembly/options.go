package assembly

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type OptionsAction struct {
	actionutils.ParentAction
}

func (this *OptionsAction) RunPost(params struct{}) {
	//this.Data["assemblys"] = []map[string]interface{}{
	//	{"id": 0, "name": "DDoS防护"},
	//	{"id": 1, "name": "云防火墙"},
	//	{"id": 2, "name": "主机防护"},
	//	{"id": 3, "name": "WEB漏洞扫描"},
	//	{"id": 4, "name": "主机漏洞扫描"},
	//}
	data := make([]map[string]interface{}, len(typeMap))
	for k, v := range typeMap {
		data[k] = map[string]interface{}{
			"id": k, "name": v,
		}
	}
	this.Data["assemblys"] = data
	this.Success()
}
