package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageSize int
	PageNo   int
}) {
	this.Data["nodeErr"] = ""
	this.Data["scans"] = make([]interface{}, 0)
	defer this.Show()
	err := webscan.InitAPIServer()
	if err != nil {
		this.Data["nodeErr"] = err.Error()
		return
	}
	if params.PageNo <= 0 {
		params.PageNo = 0
	}
	if params.PageSize <= 0 {
		params.PageSize = 999
	}
	list, err := scans_server.List(&scans.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, AdminUserId: uint64(this.AdminId())})
	if err != nil && list != nil {
		this.Data["nodeErr"] = err.Error()
		return
	}
	var scansMaps []interface{}
	if lists, ok := list["scans"]; ok {
		scansMaps = lists.([]interface{})
	}
	nessus_list, err := nessus_scans_server.History(&nessus_scans_model.HistoryReq{AdminUserId: uint64(this.AdminId())})
	if err != nil {
		this.Data["nodeErr"] = err.Error()
		return
	}
	scansMaps = append(scansMaps, nessus_list...)
	if len(nessus_list) > 0 || len(scansMaps) > 0 {
		this.Data["scans"] = scansMaps
	}
}

func (this *IndexAction) RunPost(params struct {
	PageSize int
	PageNo   int
}) {
	this.Data["scans"] = make([]interface{}, 0)
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.PageNo <= 0 {
		params.PageNo = 0
	}
	if params.PageSize <= 0 {
		params.PageSize = 999
	}
	list, err := scans_server.List(&scans.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, AdminUserId: uint64(this.AdminId())})
	//if err != nil && list != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	var scansMaps []interface{}
	if lists, ok := list["scans"]; ok {
		scansMaps = lists.([]interface{})
	}
	nessus_list, err := nessus_scans_server.History(&nessus_scans_model.HistoryReq{AdminUserId: uint64(this.AdminId())})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	scansMaps = append(scansMaps, nessus_list...)
	if len(nessus_list) > 0 || len(scansMaps) > 0 {
		this.Data["scans"] = scansMaps
	}
	this.Success()
}
