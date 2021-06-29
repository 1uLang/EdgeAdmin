package assembly

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IDCOptionsAction struct {
	actionutils.ParentAction
}

func (this *IDCOptionsAction) RunPost(params struct{}) {
	//this.Data["idcs"] = []map[string]interface{}{
	//	{"id": 0, "name": "成都IDC"},
	//	{"id": 1, "name": "杭州IDC"},
	//	{"id": 2, "name": "济南IDC"},
	//}
	data := make([]map[string]interface{}, 0)
	for k, v := range idcMap {
		data = append(data, map[string]interface{}{
			"id": k, "name": v,
		})
	}
	this.Data["idcs"] = data
	this.Success()
}
