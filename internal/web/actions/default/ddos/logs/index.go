package logs

import (
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	logs_server "github.com/1uLang/zhiannet-api/ddos/server/logs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) RunGet(params struct {
	Address string
	NodeId  uint64
	Level   int
}) {
	defer this.Show()
	this.Data["traffics"] = ""

	this.Data["ddos"] = ""
	this.Data["nodeId"] = params.NodeId
	this.Data["Address"] = params.Address
	this.Data["level"] = 1
	this.Data["page"] = ""

	if params.Level == 0 {
		params.Level = 1
	}
	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		//this.ErrorPage(err)
		return
	}
	if len(ddos) == 0 {
		//this.ErrorPage(fmt.Errorf("未配置DDoS防火墙节点"))
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
		//this.ErrorPage(err)
		return
	}

	page := this.NewPage(int64(len(list.Report)))
	this.Data["page"] = page.AsHTML()

	offset := page.Offset
	end := offset + page.Size
	if end > page.Total {
		end = page.Total
	}
	this.Data["traffics"] = list.Report[offset:end]

	this.Data["ddos"] = ddos

	this.Data["nodeId"] = params.NodeId
	this.Data["Address"] = params.Address
	this.Data["level"] = params.Level

}
