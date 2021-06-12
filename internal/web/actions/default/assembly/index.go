package assembly

import (
	"github.com/1uLang/zhiannet-api/common/server/subassemblynode"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Keyword string
}) {
	list, _, err := subassemblynode.GetNodeList()
	list, count, err := subassemblynode.GetNodeList()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	listMap := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for k, v := range list {
			listMap[k] = map[string]interface{}{
				"id":        v.Id,
				"name":      v.Name,
				"addr":      v.Addr,
				"port":      v.Port,
				"type":      v.Type,
				"idc":       v.Idc,
				"state":     v.State == 1,
				"status":    v.Status,
				"key":       v.Key,
				"secret":    v.Secret,
				"idc_name":  idcMap[v.Idc],
				"type_name": typeMap[v.Type],
			}
		}
	}
	this.Data["list"] = listMap
	page := this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	this.Show()
}
