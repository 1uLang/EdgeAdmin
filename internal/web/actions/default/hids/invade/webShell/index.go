package webShell

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

// 异常进程 相关主机
func (this *IndexAction) RunGet(params struct {
	ServerIp string
	PageNo   int
	pageSize int
}) {
	defer this.Show()

	this.Data["datas"] = nil
	this.Data["serverIp"] = params.ServerIp

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &risk.RiskSearchReq{}
	req.ServerIp = params.ServerIp
	req.PageSize = params.pageSize
	req.PageNo = params.PageNo
	req.UserName, err = this.UserName()
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取用户信息失败：%v", err)
		return
	}
	list, err := risk_server.WebShellList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取网页后门入侵威胁列表失败：%v", err)
		return
	}
	for k, v := range list.WebshellCountInfoList {

		if v["userName"] != req.UserName {
			continue
		}
		os, err := server.Info(v["serverIp"].(string), req.UserName)
		if err != nil {
			this.Data["errorMessage"] = fmt.Sprintf("获取主机信息失败：%v", err)
			return
		}
		list.WebshellCountInfoList[k]["os"] = os
	}
	this.Data["datas"] = list.WebshellCountInfoList
	this.Data["serverIp"] = params.ServerIp

}

// 异常进程 忽略
func (this *IndexAction) RunPost(params struct {
	Opt     string
	MacCode string
	RiskIds []int
	ItemIds []string
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)

	}
	req := &risk.ProcessReq{Opt: params.Opt}
	req.Req.MacCode = params.MacCode
	req.Req.RiskIds = params.RiskIds
	req.Req.ItemIds = params.ItemIds

	err = risk_server.ProcessWebShell(req)
	if err != nil {
		this.ErrorPage(err)

	}
	this.Success()
}

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) Init() {
	this.FirstMenu("index")
}

// 异常进程列表
func (this *DetailAction) RunGet(params struct {
	MacCode   string
	RiskId    string
	IsProcess bool
	Must      *actions.Must
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")
	params.Must.
		Field("riskId", params.MacCode).
		Require("请输入异常进程id")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)

	}

	info, err := risk_server.WebShellDetail(params.MacCode, params.RiskId, params.IsProcess)
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["details"] = info

	this.Show()
}

type DetailListAction struct {
	actionutils.ParentAction
}

func (this *DetailListAction) Init() {
	this.FirstMenu("index")
}

// 异常进程列表
func (this *DetailListAction) RunGet(params struct {
	Ip       string
	MacCode  string
	PageNo   int
	PageSize int
	Must     *actions.Must
}) {

	params.Must.Field("macCode", params.MacCode).Require("请输入机器码")
	params.Must.Field("ip", params.Ip).Require("请输入主机ip")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
	}
	req := &risk.DetailReq{}
	req.Req.PageSize = params.PageSize
	req.Req.PageNo = params.PageNo
	req.Req.UserName, err = this.UserName()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取用户信息失败：%v", err))
	}
	req.MacCode = params.MacCode

	//待处理
	req.Req.IsProcessed = false
	req.Req.State = 0
	list1, err := risk_server.WebShellDetailList(req)
	if err != nil {
		this.ErrorPage(err)
	}
	//已处理 - 隔离
	req.Req.IsProcessed = true
	req.Req.State = 1
	list2, err := risk_server.WebShellDetailList(req)
	if err != nil {
		this.ErrorPage(err)
	}
	//已处理 - 信任
	req.Req.IsProcessed = true
	req.Req.State = 2
	list3, err := risk_server.WebShellDetailList(req)
	if err != nil {
		this.ErrorPage(err)
	}
	//已处理 - 删除
	req.Req.IsProcessed = true
	req.Req.State = 3
	list4, err := risk_server.WebShellDetailList(req)
	if err != nil {
		this.ErrorPage(err)
	}
	//漏洞列表
	this.Data["datas1"] = list1.WebshellInfoLis
	this.Data["datas2"] = list2.WebshellInfoLis
	this.Data["datas3"] = list3.WebshellInfoLis
	this.Data["datas4"] = list4.WebshellInfoLis

	this.Data["total1"] = list1.TotalData
	this.Data["total2"] = list2.TotalData
	this.Data["total3"] = list3.TotalData
	this.Data["total4"] = list4.TotalData

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode
	//os
	os, err := server.Info(params.Ip, req.Req.UserName)
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["os"] = os["osType"]
	//最后扫描时间
	this.Data["lastUpdateTime"] = os["lastUpdateTime"]

	this.Show()
}
