package monitor

import (
	monitor_list_server "github.com/1uLang/zhiannet-api/monitor/server/monitor_list"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	Id uint64

	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择监控id")

	_, err := monitor_list_server.Del(&monitor_list_server.DelReq{Id: params.Id})
	if err != nil {
		this.ErrorPage(err)
	}
	// 记录日志
	//defer this.CreateLog(oplogs.LevelInfo, "从DNS服务商中删除域名 %d", params.DomainId)

	this.Success()
}
