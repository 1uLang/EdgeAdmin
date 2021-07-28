package databackup

import (
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			//Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeBackup)).
			//Helper(new(Helper)).
			//Helper(settingutils.NewAdvancedHelper("dbNodes")).
			//Prefix("/db").
			//Get("", new(IndexAction)).
			//GetPost("/createPopup", new(CreatePopupAction)).
			//GetPost("/updatePopup", new(UpdatePopupAction)).
			//Post("/delete", new(DeleteAction)).
			//GetPost("/cleanPopup", new(CleanPopupAction)).
			//Post("/deleteTable", new(DeleteTableAction)).
			//Post("/truncateTable", new(TruncateTableAction)).

			EndAll()
	})
}
