package hostlist

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"github.com/1uLang/zhiannet-api/zstack/server/host_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type DiskAction struct {
	actionutils.ParentAction
}

func (this *DiskAction) RunGet(params struct {
}) {

	//
	list, err := host_server.DiskList(&host.DiskListReq{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if list == nil {
		this.Data["list"] = []interface{}{}
	} else {
		lists := []maps.Map{}
		for _, v := range list.Inventories {
			memory := v.DiskSize / 1024 / 1024 / 1024
			lists = append(lists, maps.Map{
				"name":       v.Name,
				"size":       fmt.Sprintf("%v G", memory),
				"createTime": v.CreateDate,
				"uuid":       v.UUID,
			})
		}
		this.Data["list"] = lists
	}

	this.Success()
}
