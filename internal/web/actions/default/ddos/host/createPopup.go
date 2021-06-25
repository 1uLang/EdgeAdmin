package host

import (
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
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
	UserId uint64

	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Addr).
		Require("请输入ip地址").
		Field("user", params.UserId).
		Require("请选择所属用户")
	id, err := host_status_server.AddAddr(&ddos_host_ip.AddHost{
		Addr:   params.Addr,
		NodeId: params.NodeId,
		UserId: params.UserId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["ddos"] = maps.Map{
		"id":   id,
		"addr": params.Addr,
		"user": params.UserId,
	}

	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "创建高防IP %d", params.Addr)

	this.Success()
}
