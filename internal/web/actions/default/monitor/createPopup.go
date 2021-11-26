package monitor

import (
	monitor_list_server "github.com/1uLang/zhiannet-api/monitor/server/monitor_list"
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
	Addr string
	Port int

	Must *actions.Must
}) {
	params.Must.
		Field("addr", params.Addr).
		Require("请输入IP地址")

	params.Must.
		Field("port", params.Port).
		Require("请输入监控端口")

	_, err := monitor_list_server.Add(&monitor_list_server.AddReq{Addr: params.Addr, MonitorType: 1,
		Port: params.Port})
	if err != nil {
		this.ErrorPage(err)
	}
	//// 创建日志
	//defer this.CreateLog(oplogs.LevelInfo, "创建高防IP %d", params.Addr)

	this.Success()
}
