package hostlist

import (
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"github.com/1uLang/zhiannet-api/zstack/server/host_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type CandidateAction struct {
	actionutils.ParentAction
}

func (this *CandidateAction) RunGet(params struct {
	Uuid string
}) {

	//
	list, err := host_server.MigrationCandidateHost(&host.ActionReq{
		Uuid: params.Uuid,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if list == nil {
		this.Data["list"] = []interface{}{}
	} else {
		lists := []maps.Map{}
		for _, v := range list.Inventories {
			lists = append(lists, maps.Map{
				"name":       v.Name,
				"ip":         v.ManagementIp,
				"vrSkill":    v.HypervisorType,
				"createTime": v.CreateDate,
				"uuid":       v.UUID,
			})
		}
		this.Data["list"] = lists
	}

	this.Success()
}
