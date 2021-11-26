package node

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "index")
}

func (this *IndexAction) RunGet(params struct {
	NodeId int64
}) {
	nodeResp, err := this.RPC().MonitorNodeRPC().FindEnabledMonitorNode(this.AdminContext(), &pb.FindEnabledMonitorNodeRequest{MonitorNodeId: params.NodeId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	node := nodeResp.MonitorNode
	if node == nil {
		this.NotFound("monitorNode", params.NodeId)
		return
	}

	this.Data["node"] = maps.Map{
		"id":          node.Id,
		"name":        node.Name,
		"description": node.Description,
		"isOn":        node.IsOn,
	}

	this.Show()
}
