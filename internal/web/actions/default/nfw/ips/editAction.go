package ips

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/server/ips"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type EditActAction struct {
	actionutils.ParentAction
}

func (this *EditActAction) RunPost(params struct {
	Id     int64
	NodeId uint64
	Act    string
}) {
	_, err := ips.EditAction(&ips.EditActionReq{
		NodeId: params.NodeId,
		Sid:    params.Id,
		Action: params.Act,
	})
	if err != nil {
		this.ErrorPage(fmt.Errorf("修改失败"))
		return
	}
	// 记录日志
	defer this.CreateLog(oplogs.LevelInfo, "修改ips 操作%v 节点%d ID", params.Act, params.NodeId, params.Id)
	this.Success()
}
