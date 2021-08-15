package tasks

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/clusterutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(clusterutils.NewClustersHelper()).Data("teaMenu", "waf").
			Prefix("/clusters/tasks").
			GetPost("/listPopup", new(ListPopupAction)).
			GetPost("/check", new(CheckAction)).
			Post("/delete", new(DeleteAction)).
			Post("/deleteBatch", new(DeleteBatchAction)).
			EndAll()
	})
}
