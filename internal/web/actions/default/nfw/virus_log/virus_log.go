package virus_log

import (
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	request_clamav "github.com/1uLang/zhiannet-api/opnsense/request/clamav"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/clamav"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type VirusLogAction struct {
	actionutils.ParentAction
}

func (this *VirusLogAction) Init() {
	this.Nav("", "", "")
}

func (this *VirusLogAction) RunGet(params struct {
	NodeId   uint64
	PageSize int
	Page     int
}) {
	node, _, err := opnsense_server.GetOpnsenseNodeList()
	if err != nil || node == nil {
		//this.ErrorPage(err)
		//return
		node = make([]*subassemblynode.Subassemblynode, 0)
	}
	// 规则列表
	if params.NodeId == 0 && len(node) > 0 {
		params.NodeId = node[0].Id
	}
	if params.PageSize == 0 {
		params.PageSize = 20
	}
	if params.Page == 0 {
		params.Page = 1
	}

	list, err := clamav.GetLog(&clamav.LogReq{
		NodeId:   params.NodeId,
		RowCount: int(params.PageSize),
		Current:  int(params.Page),
	})
	if err != nil || list == nil {
		list = &request_clamav.LogResp{}
	}
	count := list.Total
	page := this.NewPage(int64(count))
	this.Data["page"] = page.AsHTML()
	if len(list.Rows) > 0 {
		//list.Rows[8].Line = fmt.Sprintf("Mon Sep 27 16:16:05 2021  -> Database update  faile.")

		this.Data["tableData"] = list.Rows
	} else {
		this.Data["tableData"] = make([]interface{}, 0)
	}
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	this.Show()
	//this.Success()
}
