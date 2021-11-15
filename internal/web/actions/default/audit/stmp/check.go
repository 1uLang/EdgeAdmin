package stmp

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/email"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CheckAction struct {
	actionutils.ParentAction
}

func (this *CheckAction) Init() {
	this.Nav("", "", "create")
}

func (this *CheckAction) RunPost(params struct {
	Host     string `json:"host"`
	Email    string `json:"email"`
	Port     uint   `json:"port"`
	Password string `json:"password"`
	To       string `json:"to"`

	Must *actions.Must
}) {
	params.Must.
		Field("email", params.Email).
		Require("请输入邮箱")
	if params.Host == "" {
		this.Fail("请配置邮件服务器")
	}
	if params.Port == 0 {
		this.Fail("请配置邮件端口")
	}
	if params.Password == "" {
		this.Fail("请配置密码")
	}
	if params.To == "" {
		this.Fail("请填写收件测试邮箱")
	}
	res, err := email.CheckEmail(&email.SetEmailReq{
		User: &request.UserReq{
			AdminUserId: uint64(this.AdminId()),
		},
		Host:     params.Host,
		Username: params.Email,
		Port:     params.Port,
		Password: params.Password,
		To:       params.To,
	})
	fmt.Println("测试失败", res, err)
	if err != nil || res == nil || (res != nil && res.Code != 0) {
		this.Fail("测试失败")
		return
	}

	this.Success()
}
