package risk

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DangerAccountAction struct {
	actionutils.ParentAction
}

func (this *DangerAccountAction) Init() {
	this.FirstMenu("index")
}

// 危险账号 相关主机
func (this *DangerAccountAction) RunGet(params struct {
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
	req.ServerIp = params.ServerIp
	req.PageSize = params.PageSize
	req.PageNo = params.PageNo
	req.UserName = "luobing"
	list, err := risk_server.DangerAccountList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for k, v := range list.List {
		os, err := server.Info(v["serverIp"].(string))
		if err != nil {
			this.ErrorPage(err)
		}
		list.List[k]["os"] = os
	}
	this.Data["dangerAccounts"] = list.List
	this.Data["serverIp"] = params.ServerIp
	this.Show()
}

// 危险账号 忽略
func (this *DangerAccountAction) RunPost(params struct {
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

	err = risk_server.ProcessDangerAccount(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type DangerAccountListAction struct {
	actionutils.ParentAction
}

func (this *DangerAccountListAction) Init() {
	this.FirstMenu("index")
}

// 危险账号列表
func (this *DangerAccountListAction) RunGet(params struct {
	Ip             string
	MacCode        string
	Os             string
	LastUpdateTime string
	PageNo         int
	PageSize       int
}) {

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.DetailReq{}
	req.MacCode = params.MacCode
	req.Req.PageSize = params.PageSize
	req.Req.PageNo = params.PageNo
	req.Req.UserName = "luobing"

	//待处理
	req.Req.ProcessState = 1
	list1, err := risk_server.DangerAccountDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//已处理
	req.Req.ProcessState = 2
	list2, err := risk_server.DangerAccountDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//漏洞列表
	this.Data["dangerAccount1"] = list1.DangerAccountList
	this.Data["dangerAccount2"] = list2.DangerAccountList

	this.Data["total1"] = list1.TotalData
	this.Data["total2"] = list2.TotalData

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode

	this.Data["os"] = params.Os
	//最后扫描时间
	this.Data["lastUpdateTime"] = params.LastUpdateTime

	this.Show()
}

type DangerAccountDetailAction struct {
	actionutils.ParentAction
}

func (this *DangerAccountDetailAction) Init() {
	this.FirstMenu("index")
}

// 弱口令详情
func (this *DangerAccountDetailAction) RunGet(params struct {
	MacCode      string
	RiskId       string
	ProcessState int

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	info, err := risk_server.DangerAccountDetail(params.MacCode, params.RiskId, params.ProcessState == 2)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["DangerAccountDetails"] = info

	this.Show()
}
