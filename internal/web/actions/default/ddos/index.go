package ddos

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
	Keyword  string
	NodeId   uint64
	PageNum  int
	PageSize int
}) {
	if params.NodeId == 0 {
		params.NodeId = 1
	}
	list, total, err := host_status_server.GetHostList(&ddos_host_ip.HostReq{
		NodeId:   params.NodeId,
		Addr:     params.Keyword,
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	this.Data["list"] = list
	this.Data["total"] = total
	this.Data["ddos"] = ddos
	this.Show()
	//this.Success()
}
