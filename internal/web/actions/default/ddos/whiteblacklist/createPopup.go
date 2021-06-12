package whiteblacklist

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{}) {
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Addr   string
	NodeId uint64

	Must *actions.Must
}) {
	//params.Must.
	//	Field("name", params.Addr).
	//	Require("请输入ip地址")
	//id,err := host_status_server.AddAddr(&ddos_host_ip.AddHost{
	//	Addr:   params.Addr,
	//	NodeId: params.NodeId,
	//})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	//
	//this.Data["ddos"] = maps.Map{
	//	"id":   id,
	//	"addr": params.Addr,
	//}
	//
	//// 创建日志
	//defer this.CreateLog(oplogs.LevelInfo, "创建高防IP %d", params.Addr)

	this.Success()
}
