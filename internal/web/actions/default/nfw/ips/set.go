package ips

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/server/ips"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
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
	Id     int64
	NodeId uint64
}) {
	res, err := ips.EditIps(&ips.EditIpsReq{
		NodeId: params.NodeId,
		Sid:    params.Id,
	})
	if err != nil || !res {
		this.ErrorPage(fmt.Errorf("修改失败"))
		return
	}
	//// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "修改ips规则状态 节点%v ID%d", params.NodeId, params.Id)

	this.Success()
}
