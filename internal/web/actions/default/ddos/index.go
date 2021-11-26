package ddos

import (
	"fmt"
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
	defer this.Show()

	this.Data["list"] = nil
	this.Data["total"] = 0
	this.Data["ddos"] = ""
	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	if len(ddos) == 0 {
		this.Data["errorMessage"] = fmt.Sprintf("未配置DDoS防火墙节点")
		return
	}
	if params.NodeId == 0 {
		params.NodeId = ddos[0].Id
	}
	list, total, err := host_status_server.GetHostList(&ddos_host_ip.HostReq{
		NodeId:   params.NodeId,
		Addr:     params.Keyword,
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
	})
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取DDoS防火墙列表失败：%v", err.Error())
		return
	}
	this.Data["list"] = list
	this.Data["total"] = total
	this.Data["ddos"] = ddos
}
