package invade

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type LogDeleteAction struct {
	actionutils.ParentAction
}

func (this *LogDeleteAction) Init() {
	this.FirstMenu("index")
}

// 日志异常删除 相关主机
func (this *LogDeleteAction) RunGet(params struct {
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
	list, err := risk_server.LogDeleteList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}

// 日志异常删除 忽略
func (this *LogDeleteAction) RunPost(params struct {
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

	err = risk_server.ProcessLogDelete(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type LogDeleteDetailAction struct {
	actionutils.ParentAction
}

func (this *LogDeleteDetailAction) Init() {
	this.FirstMenu("index")
}

// 日志异常删除列表
func (this *LogDeleteDetailAction) RunGet(params struct {
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

	list, err := risk_server.LogDeleteDetail(params.MacCode, params.RiskId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}

type LogDeleteDetailListAction struct {
	actionutils.ParentAction
}

func (this *LogDeleteDetailListAction) Init() {
	this.FirstMenu("index")
}

// 日志异常删除列表
func (this *LogDeleteDetailListAction) RunGet(params struct {
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

	list, err := risk_server.LogDeleteDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}
