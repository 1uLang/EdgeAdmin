package invade

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type SystemCmdAction struct {
	actionutils.ParentAction
}

func (this *SystemCmdAction) Init() {
	this.FirstMenu("index")
}

// 日志异常删除 相关主机
func (this *SystemCmdAction) RunGet(params struct {
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
	list, err := risk_server.SystemCmdList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}

// 日志异常删除 忽略
func (this *SystemCmdAction) RunPost(params struct {
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

	err = risk_server.ProcessSystemCmd(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type SystemCmdDetailAction struct {
	actionutils.ParentAction
}

func (this *SystemCmdDetailAction) Init() {
	this.FirstMenu("index")
}

// 日志异常删除列表
func (this *SystemCmdDetailAction) RunGet(params struct {
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
		Require("请输入日志异常删除id")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	list, err := risk_server.SystemCmdDetail(params.MacCode, params.RiskId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}

type SystemCmdDetailListAction struct {
	actionutils.ParentAction
}

func (this *SystemCmdDetailListAction) Init() {
	this.FirstMenu("index")
}

// 日志异常删除列表
func (this *SystemCmdDetailListAction) RunGet(params struct {
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

	list, err := risk_server.SystemCmdDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}
