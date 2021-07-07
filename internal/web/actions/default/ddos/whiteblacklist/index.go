package whiteblacklist

import (
	"fmt"
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
	defer this.Show()
	this.Data["list"] = nil
	this.Data["ddos"] = nil
	this.Data["Address"] = ""
	this.Data["nodeId"] = ""

	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	if len(ddos) == 0 {
		this.Data["errorMessage"] = "未配置DDoS防火墙节点"
		return
	}
	this.Data["ddos"] = ddos
	if params.NodeId == 0 {
		params.NodeId = ddos[0].Id
	}
	req := &black_white_list_server.BWReq{
		NodeId: params.NodeId,
		Addr:   params.Address,
	}
	this.Data["nodeId"] = req.NodeId
	list, err := black_white_list_server.GetBWList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取DDoS防火墙黑白名单列表失败：%v", err.Error())
		return
	}
	this.Data["list"] = list.Bwlist
	this.Data["Address"] = list.Address

}
