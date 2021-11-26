package logs

import (
	"fmt"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	logs_server "github.com/1uLang/zhiannet-api/ddos/server/logs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type LinkAction struct {
	actionutils.ParentAction
}

func (this *LinkAction) RunGet(params struct {
	Address string
	NodeId  uint64
	Level   int
}) {
	defer this.Success()
	if params.Level == 0 {
		params.Level = 1
	}
	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.Error(err.Error(), 400)
		return
	}
	if len(ddos) == 0 {
		this.Error("未配置DDoS防火墙节点", 400)
		return
	}
	if params.NodeId == 0 {
		params.NodeId = ddos[0].Id
	}
	req := &logs_server.LinkLogReq{
		NodeId: params.NodeId,
		Addr:   params.Address,
		Level:  params.Level,
	}
	list, err := logs_server.GetLinkLogList(req)
	if err != nil {
		this.Error(fmt.Sprintf("获取流量分析列表失败：%v", err), 400)
		return
	}

	this.Data["links"] = list.Report

	this.Data["ddos"] = ddos

	this.Data["nodeId"] = params.NodeId
	this.Data["address"] = params.Address
	this.Data["level"] = params.Level
}
