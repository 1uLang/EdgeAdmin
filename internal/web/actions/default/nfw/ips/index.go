package ips

import (
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/ips"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	NodeId   uint64
	PageNum  int
	PageSize int
}) {
	node, _, err := opnsense_server.GetOpnsenseNodeList()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 规则列表
	if params.NodeId == 0 && len(node) > 0 {
		params.NodeId = node[0].Id
	}
	list, err := ips.GetIpsList(&ips.IpsReq{
		NodeId: params.NodeId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if len(list.Rows) > 0 {
		this.Data["tableData"] = list.Rows

	}
	//this.Data["tableDataList"] = list
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	this.Show()
	//this.Success()
}
