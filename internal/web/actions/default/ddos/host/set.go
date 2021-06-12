package host

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type SetAction struct {
	actionutils.ParentAction
}

func (this *SetAction) Init() {
	this.Nav("", "", "")
}

func (this *SetAction) RunGet(params struct{}) {
	this.Show()
}

func (this *SetAction) RunPost(params struct {
	Addr   string
	Ignore bool //IP直通
	Set    int  //防护等级

	Must *actions.Must
}) {
	//params.Must.
	//	Field("name", params.Addr).
	//	Require("请输入ip地址")
	//id, err := host_status_server.AddAddr(&ddos_host_ip.AddHost{
	//	Addr:   params.Addr,
	//	//NodeId: params.Set,
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
