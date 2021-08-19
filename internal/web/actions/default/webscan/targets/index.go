package targets

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"

	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
)

//任务目标
type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageSize int
	PageNo   int

	Show int
}) {
	data := make([]interface{}, 0)
	this.Data["nodeErr"] = ""
	this.Data["targets"] = data
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取扫描节点信息失败：%v", err))
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 999
	}
	list, err := targets_server.List(&targets.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, AdminUserId: uint64(this.AdminId())})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	var targetsMaps []interface{}
	if lists, ok := list["targets"]; ok {
		targetsMaps = lists.([]interface{})
	}
	nessus_list, err := nessus_scans_server.List(&nessus_scans_model.ListReq{ AdminUserId: uint64(this.AdminId()),Targets: true})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	targetsMaps = append(targetsMaps, nessus_list...)
	if len(nessus_list) > 0 || len(targetsMaps)> 0{
		this.Data["targets"] = targetsMaps
	}
	if params.Show == 1 {
		this.Success()
	}
	this.Show()
}
func (this *IndexAction) RunPost(params struct {
	PageSize int
	PageNo   int
}) {
	data := make([]interface{}, 0)
	this.Data["nodeErr"] = ""
	this.Data["targets"] = data
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取扫描节点信息失败：%v", err))
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 999
	}
	list, err := targets_server.List(&targets.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, AdminUserId: uint64(this.AdminId())})
	//if err != nil && list != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	var targetsMaps []interface{}
	if lists, ok := list["targets"]; ok {
		targetsMaps = lists.([]interface{})
	}
	nessus_list, err := nessus_scans_server.List(&nessus_scans_model.ListReq{ AdminUserId: uint64(this.AdminId()),Targets: true})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	targetsMaps = append(targetsMaps, nessus_list...)
	if len(nessus_list) > 0 || len(targetsMaps)> 0{
		this.Data["targets"] = targetsMaps
	}
	this.Success()
}
