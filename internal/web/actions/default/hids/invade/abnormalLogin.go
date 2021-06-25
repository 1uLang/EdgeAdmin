package invade

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type AbnormalLoginAction struct {
	actionutils.ParentAction
}

func (this *AbnormalLoginAction) Init() {
	this.FirstMenu("index")
}

// 异常登录 相关主机
func (this *AbnormalLoginAction) RunGet(params struct {
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
	list, err := risk_server.AbnormalLoginList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}

// 异常登录 忽略
func (this *AbnormalLoginAction) RunPost(params struct {
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

	err = risk_server.ProcessAbnormalLogin(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type AbnormalLoginDetailAction struct {
	actionutils.ParentAction
}

func (this *AbnormalLoginDetailAction) Init() {
	this.FirstMenu("index")
}

// 异常登录列表
func (this *AbnormalLoginDetailAction) RunGet(params struct {
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
		Require("请输入异常登录id")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	list, err := risk_server.AbnormalLoginDetail(params.MacCode, params.RiskId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}

type AbnormalLoginDetailListAction struct {
	actionutils.ParentAction
}

func (this *AbnormalLoginDetailListAction) Init() {
	this.FirstMenu("index")
}

// 异常登录列表
func (this *AbnormalLoginDetailListAction) RunGet(params struct {
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

	list, err := risk_server.AbnormalLoginDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}
