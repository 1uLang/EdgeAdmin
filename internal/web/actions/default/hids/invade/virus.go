package invade

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
)

type VirusAction struct {
	actionutils.ParentAction
}

func (this *VirusAction) Init() {
	this.FirstMenu("index")
}

// 弱口令 相关主机
func (this *VirusAction) RunGet(params struct {
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
	list, err := risk_server.WeakList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Success()
}

// 弱口令 忽略
func (this *VirusAction) RunPost(params struct {
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

	err = risk_server.ProcessWeak(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type VirusDetailAction struct {
	actionutils.ParentAction
}

func (this *VirusDetailAction) Init() {
	this.FirstMenu("index")
}

// 弱口令列表
func (this *VirusDetailAction) RunGet(params struct {
	MacCode      string
	PageNo       int
	PageSize     int
	ProcessState int

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	this.Show()
	return
	//params.Must.
	//	Field("macCode", params.MacCode).
	//	Require("请输入机器码")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.DetailReq{MacCode: params.MacCode}
	req.Req.PageNo = params.PageNo
	req.Req.PageSize = params.PageSize
	req.Req.ProcessState = params.ProcessState

	list, err := risk_server.WeakDetail(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Success()
}

type VirusDetailListAction struct {
	actionutils.ParentAction
}

func (this *VirusDetailListAction) Init() {
	this.FirstMenu("index")
}

// 弱口令列表
func (this *VirusDetailListAction) RunGet(params struct {
	MacCode      string
	PageNo       int
	PageSize     int
	ProcessState int

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	this.Show()
	return
	//params.Must.
	//	Field("macCode", params.MacCode).
	//	Require("请输入机器码")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.DetailReq{MacCode: params.MacCode}
	req.Req.PageNo = params.PageNo
	req.Req.PageSize = params.PageSize
	req.Req.ProcessState = params.ProcessState

	list, err := risk_server.WeakDetail(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Success()
}
