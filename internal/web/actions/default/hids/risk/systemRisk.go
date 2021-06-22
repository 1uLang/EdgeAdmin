package risk

import (
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
	pageSize int
}) {
	this.Show()
	return
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.SearchReq{}
	req.Level = params.Level
	req.ServerIp = params.ServerIp
	req.PageSize = params.pageSize
	req.PageNo = params.PageNo
	list, err := risk_server.SystemRiskList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}

// 系统漏洞 忽略
func (this *SystemRiskAction) RunPost(params struct {
	Opt     string
	MacCode string
	RiskIds []string
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
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type SystemRiskListAction struct {
	actionutils.ParentAction
}

func (this *SystemRiskListAction) Init() {
	this.FirstMenu("index")
}

// 系统漏洞列表
func (this *SystemRiskListAction) RunGet(params struct {
	MacCode string
	RiskId  string

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	this.Show()
	return
	//params.Must.
	//	Field("macCode", params.MacCode).
	//	Require("请输入机器码")
	//
	//params.Must.
	//	Field("riskId", params.RiskId).
	//	Require("请输入系统漏洞id")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	list, err := risk_server.SystemRiskDetail(params.MacCode, params.RiskId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}
