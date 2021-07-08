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
	defer this.Show()

	this.Data["attacks"] = nil
	this.Data["ddos"] = nil
	this.Data["nodeId"] = params.NodeId
	//2006-01-02
	this.Data["startTime"] = params.StartTime
	this.Data["endTime"] = params.EndTime
	this.Data["address"] = params.Address
	this.Data["attackType"] = params.AttackType
	this.Data["status"] = params.Status

	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	if len(ddos) == 0 {
		this.Data["errorMessage"] = "未配置DDoS防火墙节点"
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
			this.Data["errorMessage"] = "起始时间参数错误"
			return
		}
		eT, err := time.Parse("2006-01-02", params.EndTime)
		if err != nil {
			this.Data["errorMessage"] = "结束时间参数错误"
			return
		}
		req.StartTime = sT
		req.EndTime = eT
	}

	list, err := logs_server.GetAttackLogList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取统计日志列表失败：%v",err)
		return
	}
	this.Data["attacks"] = list.Report
	this.Data["ddos"] = ddos
	this.Data["nodeId"] = req.NodeId
	//2006-01-02
	if len(list.StartDate) > 10{
		this.Data["startTime"] = list.StartDate[:10]
	}else{
		this.Data["startTime"] = list.StartDate
	}
	if len(list.EndDate) > 10{
		this.Data["endTime"] = list.EndDate[:10]
	}else{
		this.Data["endTime"] = list.EndDate
	}
	this.Data["address"] = list.Address
	this.Data["attackType"] = list.CurFlags
	this.Data["status"] = list.CurStatus

}
