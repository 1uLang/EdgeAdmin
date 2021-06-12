package whiteblacklist

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DelAction struct {
	actionutils.ParentAction
}

func (this *DelAction) Init() {
	this.Nav("", "", "")
}

func (this *DelAction) RunGet(params struct{}) {
	this.Show()
}

func (this *DelAction) RunPost(params struct {
	Addr   string
	Ignore bool //IP直通
	Del    int  //防护等级

	Must *actions.Must
}) {
	//params.Must.
	//	Field("name", params.Addr).
	//	Require("请输入ip地址")
	//id, err := host_status_server.AddAddr(&ddos_host_ip.AddHost{
	//	Addr:   params.Addr,
	//	//NodeId: params.Del,
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
