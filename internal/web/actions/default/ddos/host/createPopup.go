package host

import (
	"fmt"
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

func (this *CreatePopupAction) RunGet(params struct {
	NodeId uint64
	Must   *actions.Must
}) {
	params.Must.Field("NodeId", params.NodeId).Require("请选择DDoS防火墙节点")

	defer this.Show()
	this.Data["ddos"] = ""
	this.Data["node"] = ""

	//ddos节点
	ddos, err := host_status_server.GetDDoSNodeInfo(params.NodeId)
	if err != nil {
		this.Data["messageMessage"] = fmt.Sprintf("获取DDoS防火墙节点信息失败：%v", err)
		return
	}
	this.Data["ddos"] = ddos.Name + "-" + ddos.Addr
	this.Data["node"] = params.NodeId
}

func (this *CreatePopupAction) RunPost(params struct {
	Addr    string
	NodeId  uint64
	User_id uint64
	Remark  string
	Must    *actions.Must
	CSRF    *actionutils.CSRF
}) {
	params.Must.
		Field("addr", params.Addr).
		Require("请输入高防IP").
		Match("^(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\d|1\\d{2}"+
			"|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\"+
			"d|1\\d{2}|2[0-4]\\d|25[0-5])", "请输入正确的地址").
		//Field("user_id", params.User_id).
		//Require("请选择所属用户").
		Field("nodeId", params.NodeId).
		Require("请选择所属DDoS防火墙节点")

	if params.NodeId == 0 {
		this.Fail("请选择所属DDoS防火墙节点")
	}
	//if params.User_id == 0 {
	//	this.Fail("请选择所属用户")
	//}
	id, err := host_status_server.AddAddr(&ddos_host_ip.AddHost{
		Addr:   params.Addr,
		NodeId: params.NodeId,
		//UserId: params.User_id,
		Remark: params.Remark,
	})
	if err != nil {
		this.ErrorPage(fmt.Errorf("添加主机失败：%v", err))
		return
	}

	this.Data["ddos"] = maps.Map{
		"id":   id,
		"addr": params.Addr,
		//"user": params.User_id,
	}
	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "创建高防IP %d", params.Addr)
	this.Success()
}
