package hostlist

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"github.com/1uLang/zhiannet-api/zstack/server/host_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type NetworkAction struct {
	actionutils.ParentAction
}

func (this *NetworkAction) RunGet(params struct {
}) {

	//
	list, err := host_server.NetworkList(&host.NetworkListReq{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if list == nil {
		this.Data["list"] = []interface{}{}
	} else {
		this.Data["listlll"] = list.Inventories

		lists := maps.Map{
			"public":  []interface{}{},
			"private": []interface{}{},
		}
		public := []interface{}{}
		private := []interface{}{}
		for _, v := range list.Inventories {
			if v.Category == "Public" {
				public = append(public, maps.Map{
					"name":       v.Name,
					"netType":    v.Type,
					"ip":         v.IPRanges[0].NetworkCidr,
					"ipType":     fmt.Sprintf("IPv%v", v.IPVersion),
					"ipUse":      fmt.Sprintf("%v - %v", v.IPRanges[0].StartIP, v.IPRanges[0].EndIP),
					"createTime": v.CreateDate,
					"uuid":       v.UUID,
				})
			} else if v.Category == "Private" {
				private = append(private, maps.Map{
					"name":       v.Name,
					"netType":    v.Type,
					"ip":         v.IPRanges[0].NetworkCidr,
					"ipType":     fmt.Sprintf("IPv%v", v.IPVersion),
					"ipUse":      fmt.Sprintf("%v - %v", v.IPRanges[0].StartIP, v.IPRanges[0].EndIP),
					"createTime": v.CreateDate,
					"uuid":       v.UUID,
				})
			}

		}
		lists = maps.Map{
			"public":  public,
			"private": private,
		}
		this.Data["list"] = lists
	}

	this.Success()
}
