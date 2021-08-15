package regions

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Data("teaMenu", "regions").
			Data("teaSubMenu", "region").
			Prefix("/clusters/regions").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/updatePopup", new(UpdatePopupAction)).
			Post("/delete", new(DeleteAction)).
			Post("/sort", new(SortAction)).
			GetPost("/selectPopup", new(SelectPopupAction)).
			GetPost("/prices", new(PricesAction)).
			GetPost("/updatePricePopup", new(UpdatePricePopupAction)).
			EndAll()
	})
}
