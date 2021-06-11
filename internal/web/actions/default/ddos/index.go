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
	Keyword string
	NodeId  uint64
}) {
	list, _, err := host_status_server.GetHostList(&ddos_host_ip.HostReq{
		NodeId: params.NodeId,
		Addr:   params.Keyword,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["list"] = list
	this.Show()
	//this.Success()
}
