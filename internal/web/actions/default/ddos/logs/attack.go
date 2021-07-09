package logs

import (
	"fmt"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	logs_server "github.com/1uLang/zhiannet-api/ddos/server/logs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"time"
)

type AttacksAction struct {
	actionutils.ParentAction
}

func (this *AttacksAction) RunGet(params struct {
	Address    string
	NodeId     uint64
	StartTime  string
	EndTime    string
	AttackType string
	Status     int
}) {
	defer this.Success()

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
		this.ErrorPage(err)
		return
	}
	if len(ddos) == 0 {
		this.ErrorPage(fmt.Errorf("未配置DDoS防火墙节点"))
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

	fmt.Println(req.StartTime,req.EndTime)
	list, err := logs_server.GetAttackLogList(req)
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取统计日志列表失败：%v", err))
		return
	}
	this.Data["attacks"] = list.Report
	this.Data["ddos"] = ddos
	this.Data["nodeId"] = params.NodeId

	this.Data["startTime"] = params.StartTime

	this.Data["endTime"] = params.EndTime

	this.Data["address"] = params.Address
	this.Data["attackType"] = params.AttackType
	this.Data["status"] = params.Status

}
