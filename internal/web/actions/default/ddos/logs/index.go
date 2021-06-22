package logs

import (
	"fmt"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	logs_server "github.com/1uLang/zhiannet-api/ddos/server/logs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) RunGet(params struct {
	Address    string
	NodeId     uint64
	StartTime  string
	EndTime    string
	AttackType string
	Status     int
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
	req := &logs_server.AttackLogReq{
		NodeId:     params.NodeId,
		Addr:       params.Address,
		AttackType: params.AttackType,
		Status:     params.Status,
	}
	if params.StartTime != "" && params.EndTime != "" {
		sT, err := time.Parse("2006-01-02", params.StartTime)
		if err != nil {
			this.ErrorPage(fmt.Errorf("起始时间参数错误"))
			return
		}
		eT, err := time.Parse("2006-01-02", params.EndTime)
		if err != nil {
			this.ErrorPage(fmt.Errorf("结束时间参数错误"))
			return
		}
		req.StartTime = sT
		req.EndTime = eT
	}

	list, err := logs_server.GetAttackLogList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["attacks"] = list.Report
	this.Data["ddos"] = ddos
	this.Data["nodeId"] = req.NodeId
	//2006-01-02
	this.Data["startTime"] = list.StartDate[:10]
	this.Data["endTime"] = list.EndDate[:10]
	this.Data["address"] = list.Address
	this.Data["attackType"] = list.CurFlags
	this.Data["status"] = list.CurStatus
	this.Show()
}
