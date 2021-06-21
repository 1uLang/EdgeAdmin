package nfw

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Keyword string
	NodeId  uint64
}) {
	//if params.NodeId == 0 {
	//	params.NodeId = 1
	//}
	//list, total, err := host_status_server.GetHostList(&ddos_host_ip.HostReq{
	//	NodeId:   params.NodeId,
	//	Addr:     params.Keyword,
	//	PageSize: params.PageSize,
	//	PageNum:  params.PageNum,
	//})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	////ddos节点
	//ddos, _, err := host_status_server.GetDdosNodeList()
	//this.Data["list"] = list
	//this.Data["total"] = total
	//this.Data["ddos"] = ddos
	this.Show()
	//this.Success()
}
