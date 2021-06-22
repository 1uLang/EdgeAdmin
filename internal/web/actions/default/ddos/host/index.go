package host

import (
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
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
	list, total, err := host_status_server.GetHostList(&ddos_host_ip.HostReq{
		NodeId:   params.NodeId,
		Addr:     params.Address,
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["list"] = list
	this.Data["total"] = total
	this.Data["ddos"] = ddos
	this.Data["nodeId"] = params.NodeId
	this.Data["address"] = params.Address
	this.Show()
	//this.Success()
}
