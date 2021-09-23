package platformbackup

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/server/platform_backup_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type BackupAction struct {
	actionutils.ParentAction
}

func (this *BackupAction) RunPost(params struct {
}) {
	_, err := platform_backup_server.Backup()
	if err != nil {
		this.ErrorPage(fmt.Errorf("备份失败"))
		return
	}
	//defer this.CreateLog(oplogs.LevelInfo, "从acl 中删除配置 节点 %d %d", params.NodeId, params.Id)
	this.Success()
}
