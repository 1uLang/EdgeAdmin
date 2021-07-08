package ips

import (
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	req_ips "github.com/1uLang/zhiannet-api/opnsense/request/ips"
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
	list, err := ips.GetIpsList(&ips.IpsReq{
		NodeId: params.NodeId,
	})
	if err != nil || list == nil {
		list = &req_ips.IpsListResp{}
	}
	count := list.Total
	page := this.NewPage(int64(count))
	this.Data["page"] = page.AsHTML()

	list, err = ips.GetIpsList(&ips.IpsReq{
		NodeId: params.NodeId,
		PageSize: int(page.Size),
		PageNum: int(page.Current),
	})
	if err != nil || list == nil {
		list = &req_ips.IpsListResp{}
	}
	if list len(list.Rows) > 0 {
		this.Data["tableData"] = list.Rows
	} else {
		this.Data["tableData"] = make([]interface{}, 0)
	}
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	this.Show()
	//this.Success()
}
