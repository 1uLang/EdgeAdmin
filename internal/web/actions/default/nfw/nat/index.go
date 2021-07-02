package nat

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	req_nat "github.com/1uLang/zhiannet-api/opnsense/request/nat"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/nat"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	NodeId uint64
}) {
	node, _, err := opnsense_server.GetOpnsenseNodeList()
	if err != nil || node == nil {
		//this.ErrorPage(err)
		//return
		node = make([]*subassemblynode.Subassemblynode, 0)
	}
	//nat 规则列表
	if params.NodeId == 0 && len(node) > 0 {
		params.NodeId = node[0].Id
	}
	list, err := nat.GetNat1To1List(&nat.ListReq{
		NodeId: params.NodeId,
	})
	if err != nil || list == nil {
		//this.ErrorPage(err)
		//return
		list = make([]*req_nat.Nat1To1ListResp, 0)
	}
	this.Data["tableDataList"] = list
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	fmt.Println("list", list)
	this.Show()
	//this.Success()
}
