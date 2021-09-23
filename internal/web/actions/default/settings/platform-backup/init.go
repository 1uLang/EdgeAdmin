package platformbackup

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/settingutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeSetting)).
			Helper(settingutils.NewHelper("platform-backup")).
			Prefix("/settings/platform-backup").
			Get("", new(IndexAction)).
			Prefix("/settings/platform-backup/backup").
			GetPost("", new(BackupAction)).
			Prefix("/settings/platform-backup/delete").
			GetPost("", new(DeleteAction)).
			Prefix("/settings/platform-backup/clean").
			GetPost("", new(Delete30DayAction)).
			Prefix("/settings/platform-backup/recover").
			GetPost("", new(RecoverAction)).
			EndAll()
	})
}
