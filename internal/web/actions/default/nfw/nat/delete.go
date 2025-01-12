package nat

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/server/nat"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	NodeId uint64
	Id     string

	Must *actions.Must
}) {
	params.Must.Field("id", params.Id).Require("请输入规则id").
		Field("nodeId", params.NodeId).Require("请选择当前DDoS防护节点")

	res, err := nat.DelNat1To1(&nat.DelNat1To1Req{
		NodeId: params.NodeId,
		Id:     params.Id,
	})
	if err != nil || !res {
		this.ErrorPage(fmt.Errorf("删除失败"))
		return
	}
	// 记录日志
	defer this.CreateLog(oplogs.LevelInfo, "从nat 1:1 中删除配置 节点%d %d", params.NodeId, params.Id)
	this.Success()
}
