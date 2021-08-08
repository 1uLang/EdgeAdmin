package settingutils

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
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
	action.Data["teaMenu"] = "server"

	// 标签栏
	tabbar := actionutils.NewTabbar()
	var session = action.Session()
	var adminId = session.GetInt64("adminId")
	if configloaders.AllowModule(adminId, configloaders.AdminModuleCodeSetting) {
		tabbar.Add("Web服务", "", "/settings/server", "", this.tab == "server")
		tabbar.Add("管理界面设置", "", "/settings/ui", "", this.tab == "ui")
		tabbar.Add("用户界面设置", "", "/settings/user-ui", "", this.tab == "userUI")
		tabbar.Add("访问设置", "", "/settings/security", "", this.tab == "security")
		tabbar.Add("IP库", "", "/settings/ip-library", "", this.tab == "ipLibrary")
	}
	tabbar.Add("个人资料", "", "/settings/profile", "", this.tab == "profile")
	tabbar.Add("登录设置", "", "/settings/login", "", this.tab == "login")
	tabbar.Add("安全策略", "", "/settings/strategy", "", this.tab == "strategy")
	actionutils.SetTabbar(actionPtr, tabbar)

	return
}
