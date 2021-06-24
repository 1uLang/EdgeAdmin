package risk

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
)

type WeakAction struct {
	actionutils.ParentAction
}

func (this *WeakAction) Init() {
	this.FirstMenu("index")
}

// 弱口令 相关主机
func (this *WeakAction) RunGet(params struct {
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
	list, err := risk_server.WeakList(req)
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
	this.Data["weaks"] = list.List
	this.Data["serverIp"] = params.ServerIp
	this.Show()
}

// 弱口令 忽略
func (this *WeakAction) RunPost(params struct {
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
	err = risk_server.ProcessWeak(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type WeakListAction struct {
	actionutils.ParentAction
}

func (this *WeakListAction) Init() {
	this.FirstMenu("index")
}

// 弱口令列表
func (this *WeakListAction) RunGet(params struct {
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
	list1, err := risk_server.WeakDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//已处理
	req.Req.ProcessState = 2
	list2, err := risk_server.WeakDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//漏洞列表
	this.Data["weak1"] = list1.WeakInfoList
	this.Data["weak2"] = list2.WeakInfoList

	this.Data["total1"] = list1.TotalData
	this.Data["total2"] = list2.TotalData

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode

	this.Data["os"] = params.Os
	//最后扫描时间
	this.Data["lastUpdateTime"] = params.LastUpdateTime

	this.Show()
}
