package logs

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/server/logs"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	NodeId uint64
}) {

	res, err := logs.ClearLogs(&logs.NodeReq{
		NodeId: params.NodeId,
	})
	if err != nil || !res {
		this.ErrorPage(fmt.Errorf("修改失败"))
		return
	}
	// 记录日志
	defer this.CreateLog(oplogs.LevelInfo, "云防火墙 清除统计日志 节点%d", params.NodeId)

	this.Success()
}
