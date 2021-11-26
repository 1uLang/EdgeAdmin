package monitor

import (
	monitor_list_server "github.com/1uLang/zhiannet-api/monitor/server/monitor_list"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "")
}

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}

func (this *CreateAction) RunPost(params struct {
	Addr string
	Code int

	Must *actions.Must
}) {
	params.Must.
		Field("addr", params.Addr).
		Require("请输入IP地址")

	params.Must.
		Field("code", params.Code).
		Require("请输入HTTP状态码")

	_, err := monitor_list_server.Add(&monitor_list_server.AddReq{Addr: params.Addr, MonitorType: 2, Code: params.Code})
	if err != nil {
		this.ErrorPage(err)
	}
	//// 创建日志
	//defer this.CreateLog(oplogs.LevelInfo, "创建高防IP %d", params.Addr)

	this.Success()
}
