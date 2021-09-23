package platformbackup

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/server/platform_backup_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	Id int64
}) {
	err := platform_backup_server.Delete(params.Id)
	if err != nil {
		this.ErrorPage(fmt.Errorf("删除失败"))
		return
	}
	//defer this.CreateLog(oplogs.LevelInfo, "从acl 中删除配置 节点 %d %d", params.NodeId, params.Id)
	this.Success()
}

type Delete30DayAction struct {
	actionutils.ParentAction
}

func (this *Delete30DayAction) RunPost(params struct {
}) {
	err := platform_backup_server.Clean30DayData()
	if err != nil {
		this.ErrorPage(fmt.Errorf("删除失败"))
		return
	}
	//defer this.CreateLog(oplogs.LevelInfo, "从acl 中删除配置 节点 %d %d", params.NodeId, params.Id)
	this.Success()
}
