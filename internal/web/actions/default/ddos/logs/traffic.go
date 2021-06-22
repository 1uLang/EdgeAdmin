package logs

import (
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	logs_server "github.com/1uLang/zhiannet-api/ddos/server/logs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type TrafficAction struct {
	actionutils.ParentAction
}

func (this *TrafficAction) RunGet(params struct {
	Address string
	NodeId  uint64
	Level   int
}) {
	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.NodeId == 0 {
		params.NodeId = ddos[0].Id
	}

	//流量分析

	req := &logs_server.TrafficLogReq{
		NodeId: params.NodeId,
		Addr:   params.Address,
		Level:  params.Level,
	}
	list, err := logs_server.GetTrafficLogList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["traffics"] = list.Report

	this.Data["ddos"] = ddos

	this.Data["nodeId"] = params.NodeId
	this.Data["address"] = params.Address
	this.Data["level"] = params.Level
	this.Success()
}
