package nat

import (
	"fmt"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server/nat"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
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
	Id     string
	NodeId uint64
	Status string //状态 0禁用 1启用

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("ID不能为空")
	res, err := opnsense_server.StartUpNat1To1(&opnsense_server.StartNat1To1Req{
		Id:     params.Id,
		NodeId: params.NodeId,
	})
	if err != nil || !res {
		this.ErrorPage(fmt.Errorf("操作失败"))
		return
	}

	//
	//// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "修改云防火墙状态 节点=%d id=%d", params.NodeId, params.Id)

	this.Success()
}
