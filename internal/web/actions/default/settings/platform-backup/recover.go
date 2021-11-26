package platformbackup

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/server/platform_backup_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type RecoverAction struct {
	actionutils.ParentAction
}

func (this *RecoverAction) RunPost(params struct {
	Id int64
}) {
	err := platform_backup_server.Recovery(params.Id)
	if err != nil {
		this.ErrorPage(fmt.Errorf("恢复失败"))
		return
	}
	//defer this.CreateLog(oplogs.LevelInfo, "从acl 中删除配置 节点 %d %d", params.NodeId, params.Id)
	this.Success()
}
