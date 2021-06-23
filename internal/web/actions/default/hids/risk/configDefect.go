package risk

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type ConfigDefectAction struct {
	actionutils.ParentAction
}

func (this *ConfigDefectAction) Init() {
	this.FirstMenu("index")
}

// 配置缺陷 相关主机
func (this *ConfigDefectAction) RunGet(params struct {
	Level    int
	ServerIp string
	PageNo   int
	pageSize int
}) {

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
	list, err := risk_server.ConfigDefectList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}

// 配置缺陷 忽略
func (this *ConfigDefectAction) RunPost(params struct {
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

	err = risk_server.ProcessConfigDefect(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type ConfigDefectListAction struct {
	actionutils.ParentAction
}

func (this *ConfigDefectListAction) Init() {
	this.FirstMenu("index")
}

// 配置缺陷列表
func (this *ConfigDefectListAction) RunGet(params struct {
	MacCode      string
	PageNo       int
	PageSize     int
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
	req := &risk.DetailReq{MacCode: params.MacCode}
	req.Req.PageNo = params.PageNo
	req.Req.PageSize = params.PageSize
	req.Req.ProcessState = params.ProcessState

	list, err := risk_server.ConfigDefectDetail(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list

	this.Show()
}
