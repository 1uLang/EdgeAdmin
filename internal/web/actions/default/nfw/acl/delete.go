package acl

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/server/acl"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	NodeId uint64
	Id     string
}) {
	res, err := acl.DelAcl(&acl.DelAclReq{
		NodeId: params.NodeId,
		ID:     params.Id,
	})
	if err != nil || !res {
		this.ErrorPage(fmt.Errorf("删除失败"))
		return
	}
	defer this.CreateLog(oplogs.LevelInfo, "从acl 中删除配置 节点 %d %d", params.NodeId, params.Id)
	this.Success()
}
