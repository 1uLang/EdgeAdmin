package invade

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type AbnormalAccountAction struct {
	actionutils.ParentAction
}

func (this *AbnormalAccountAction) Init() {
	this.FirstMenu("index")
}

// 异常账号 相关主机
func (this *AbnormalAccountAction) RunGet(params struct {
	ServerIp string
	PageNo   int
	pageSize int
}) {

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.RiskSearchReq{}
	req.ServerIp = params.ServerIp
	req.PageSize = params.pageSize
	req.PageNo = params.PageNo
	list, err := risk_server.AbnormalAccountList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}

// 异常账号 忽略
func (this *AbnormalAccountAction) RunPost(params struct {
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

	err = risk_server.ProcessAbnormalAccount(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type AbnormalAccountDetailAction struct {
	actionutils.ParentAction
}

func (this *AbnormalAccountDetailAction) Init() {
	this.FirstMenu("index")
}

// 异常账号列表
func (this *AbnormalAccountDetailAction) RunGet(params struct {
	MacCode string
	RiskId  string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")
	params.Must.
		Field("riskId", params.MacCode).
		Require("请输入异常账号id")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	list, err := risk_server.AbnormalAccountDetail(params.MacCode, params.RiskId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}

type AbnormalAccountDetailListAction struct {
	actionutils.ParentAction
}

func (this *AbnormalAccountDetailListAction) Init() {
	this.FirstMenu("index")
}

// 异常账号列表
func (this *AbnormalAccountDetailListAction) RunGet(params struct {
	MacCode  string
	PageNo   int
	PageSize int
	State    int

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.DetailReq{MacCode: params.MacCode}
	req.Req.PageNo = params.PageNo
	req.Req.PageSize = params.PageSize
	req.Req.State = params.State

	list, err := risk_server.AbnormalAccountDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}
