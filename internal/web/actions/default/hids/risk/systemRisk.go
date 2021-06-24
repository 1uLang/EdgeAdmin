package risk

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
)

type SystemRiskAction struct {
	actionutils.ParentAction
}

func (this *SystemRiskAction) Init() {
	this.FirstMenu("index")
}

// 系统漏洞 相关主机
func (this *SystemRiskAction) RunGet(params struct {
	Level    int
	ServerIp string
	PageNo   int
	PageSize int
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.SearchReq{}
	req.Level = params.Level
	req.ServerIp = params.ServerIp
	req.PageSize = params.PageSize
	req.PageNo = params.PageNo
	req.UserName = "luobing"

	//系统漏洞数汇总
	risk, err := risk_server.SystemDistributed(req)
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["risks"] = risk.List
	this.Data["serverIp"] = params.ServerIp
	this.Data["level"] = params.Level
	this.Show()
}

// 系统漏洞 忽略
func (this *SystemRiskAction) RunPost(params struct {
	Opt     string
	MacCode string
	RiskIds []int
	ItemIds []string
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.ProcessReq{Opt: params.Opt}
	req.Req.MacCode = params.MacCode
	req.Req.RiskIds = params.RiskIds
	req.Req.ItemIds = params.ItemIds

	err = risk_server.ProcessRisk(req)
	fmt.Println(err)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
