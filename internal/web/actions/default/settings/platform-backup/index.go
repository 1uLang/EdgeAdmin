package platformbackup

import (
	"github.com/1uLang/zhiannet-api/common/model/platform_backup"
	"github.com/1uLang/zhiannet-api/common/server/platform_backup_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	PageSize int
	Page     int
}) {

	if params.PageSize == 0 {
		params.PageSize = 20
	}
	if params.Page == 0 {
		params.Page = 1
	}
	list, total, err := platform_backup_server.GetBackupList(&platform_backup.ListReq{
		PageSize: params.PageSize,
		PageNum:  params.Page,
	})
	if err != nil || list == nil {
		list = make([]*platform_backup.PlatformBackup, 0)
	}
	count := total
	page := this.NewPage(int64(count))
	this.Data["page"] = page.AsHTML()
	this.Data["tableData"] = list
	this.Show()
}
