package assembly

import (
	subassemblynode_server "github.com/1uLang/zhiannet-api/common/server/subassemblynode"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	NodeId uint64
}) {
	defer this.CreateLogInfo("删除节点 %d", params.NodeId)

	// TODO 检查用户是否有未完成的业务
	err := subassemblynode_server.Del(params.NodeId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
