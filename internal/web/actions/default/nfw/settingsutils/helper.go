package settingsutils

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type Helper struct {
	tab string
}

func NewHelper(tab string) *Helper {
	return &Helper{
		tab: tab,
	}
}

func (this *Helper) BeforeAction(actionPtr actions.ActionWrapper) (goNext bool) {
	goNext = true

	action := actionPtr.Object()

	// 左侧菜单
	action.Data["teaMenu"] = "virus"
	// 标签栏
	tabbar := actionutils.NewTabbar()
	tabbar.Add("病毒库信息", "", "/nfw/virus", "", this.tab == "virus")
	tabbar.Add("更新日志", "", "/nfw/virusLog", "", this.tab == "virusLog")
	actionutils.SetTabbar(actionPtr, tabbar)

	return
}
