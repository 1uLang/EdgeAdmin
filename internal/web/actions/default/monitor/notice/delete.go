package notice

import (
	monitor_notice_server "github.com/1uLang/zhiannet-api/monitor/server/monitor_notice"
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
	params.Must.Field("od", params.Id).Require("请选择监控消息")

	_, err := monitor_notice_server.Del(&monitor_notice_server.DelReq{Id: params.Id})
	if err != nil {
		this.ErrorPage(err)
	}
	// 记录日志
	//defer this.CreateLog(oplogs.LevelInfo, "从DNS服务商中删除域名 %d", params.DomainId)

	this.Success()
}
