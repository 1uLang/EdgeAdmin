package whiteblacklist

import (
	black_white_list_server "github.com/1uLang/zhiannet-api/ddos/server/black_white_list"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Address  string
	NodeId   uint64
	PageNum  int
	PageSize int
}) {

	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if len(ddos) > 0 && params.NodeId == 0 {
		params.NodeId = ddos[0].Id
	}
	req := &black_white_list_server.BWReq{
		NodeId: params.NodeId,
		Addr:   params.Address,
	}
	list, err := black_white_list_server.GetBWList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["list"] = list.Bwlist
	this.Data["ddos"] = ddos
	this.Data["Address"] = list.Address
	this.Data["nodeId"] = req.NodeId
	this.Show()
}
