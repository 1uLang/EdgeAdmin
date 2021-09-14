package hostlist

import (
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"github.com/1uLang/zhiannet-api/zstack/server/host_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type ActivityAction struct {
	actionutils.ParentAction
}

func (this *ActivityAction) RunPost(params struct {
	Event    string
	Uuid     string //云主机Ip
	HostUuid string //物理机ID
	SpecUuid string //规格ID
}) {

	switch params.Event {
	case "start":
		host_server.StartHost(&host.ActionReq{
			Uuid: params.Uuid,
		})
	case "stop":
		host_server.StopHost(&host.ActionReq{
			Uuid: params.Uuid,
		})
	case "suspend":
		host_server.Suspend(&host.SuspendReq{
			Uuid: params.Uuid,
		})
	case "unsuspend":
		host_server.UnSuspend(&host.SuspendReq{
			Uuid: params.Uuid,
		})
	case "restart":
		host_server.RestartHost(&host.ActionReq{
			Uuid: params.Uuid,
		})

	case "migration": //迁移
		host_server.MigrationHost(&host.ActionReq{
			Uuid:     params.Uuid,
			HostUUid: params.HostUuid,
		})
	case "delete":
		host_server.DeleteHost(&host.ActionReq{
			Uuid: params.Uuid,
		})
	case "spec":
		host_server.UpdateSpec(&host.UpdateSpecReq{
			SpecUUid: params.SpecUuid,
			HostUUid: params.Uuid,
		})
	}

	this.Success()
}
