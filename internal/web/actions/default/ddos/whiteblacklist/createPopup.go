package whiteblacklist

import (
	black_white_list_server "github.com/1uLang/zhiannet-api/ddos/server/black_white_list"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{ NodeId uint64 }) {

	//ddos节点
	ddos, err := host_status_server.GetDDoSNodeInfo(params.NodeId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["ddos"] = ddos.Name + "-ddos-" + ddos.Addr
	this.Data["node"] = params.NodeId
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Addr   string
	NodeId uint64
	White  string
	Must   *actions.Must
}) {
	params.Must.
		Field("Addr", params.Addr).
		Require("请输入ip地址")

	_, err := black_white_list_server.AddBW(&black_white_list_server.EditBWReq{
		Addr:   []string{params.Addr},
		NodeId: params.NodeId,
		White:  params.White == "白名单",
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
