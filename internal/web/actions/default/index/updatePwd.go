package index

import (
	edge_admins_server "github.com/1uLang/zhiannet-api/common/server/edge_admins_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/dlclark/regexp2"
	"github.com/iwind/TeaGo/actions"
	stringutil "github.com/iwind/TeaGo/utils/string"
)

type UpdatePwdAction struct {
	actionutils.ParentAction
}

// 首页（登录页）

var TokenSalt1 = stringutil.Rand(32)

func (this *UpdatePwdAction) RunGet(params struct {
	From string
	Code string
	Auth *helpers.UserShouldAuth
}) {
	if params.Code == "" {
		this.RedirectURL("/")
		return
	}
	_, err := decode(params.Code)
	if err != nil {
		this.RedirectURL("/")
		return
	}

	this.Show()
}

// RunPost 提交
func (this *UpdatePwdAction) RunPost(params struct {
	Password        string
	ConfirmPassword string
	Code            string
	Must            *actions.Must
	Auth            *helpers.UserShouldAuth
	CSRF            *actionutils.CSRF
}) {
	if params.Code == "" {
		this.FailField("refresh", "页面信息已过期，请刷新后重试")
	}
	adminId, err := decode(params.Code)
	if err != nil {
		this.FailField("refresh", "页面信息已过期，请刷新后重试")
	}

	if params.Password != params.ConfirmPassword {
		this.FailField("password", "两次输入的密码不一致")
	}
	if params.Password == stringutil.Md5("") {
		this.FailField("password", "请输入密码")
	}
	reg, err := regexp2.Compile(
		`^(?![A-z0-9]+$)(?=.[^%&',;=?$\x22])(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).{8,30}$`, 0)
	if err != nil {
		this.FailField("pass1", "密码格式不正确")
	}
	if match, err := reg.FindStringMatch(params.ConfirmPassword); err != nil || match == nil {
		this.FailField("pass1", "密码格式不正确")
	}
	res, err := edge_admins_server.UpdatePwd(uint64(adminId), stringutil.Md5(params.Password))
	if err != nil || res == 0 {
		this.Fail("修改密码失败")
	}
	//清除session
	params.Auth.Logout()

	this.Success()
}
