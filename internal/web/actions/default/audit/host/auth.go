package host

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server"
	"github.com/1uLang/zhiannet-api/audit/server/audit_host"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type AuthAction struct {
	actionutils.ParentAction
}

func (this *AuthAction) Init() {
	this.Nav("", "", "")
}

func (this *AuthAction) RunGet(params struct {
	Id uint64
}) {
	list, err := audit_host.GetAuthEmail(&server.AuthReq{
		Id: params.Id,
		User: &request.UserReq{
			AdminUserId: uint64(this.AdminId()),
		},
	})
	var email string
	if err != nil || list == nil {

	} else {
		for _, v := range list.Data.Email {
			if e := strings.TrimSpace(v); e != "" {
				email = fmt.Sprintf("%v\n%v", email, e)
			}

		}
	}
	email = strings.TrimSpace(email)
	this.Data["authValue"] = email
	this.Data["id"] = params.Id
	this.Show()
}

func (this *AuthAction) RunPost(params struct {
	Email string
	Id    uint64

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("email", params.Email).
		Require("请输入邮箱")

	emails := []string{}
	emails = strings.Split(params.Email, "\n")

	for k, v := range emails {
		emails[k] = strings.TrimSpace(v)
	}

	res, err := audit_host.AuthHost(&server.AuthReq{
		User: &request.UserReq{
			AdminUserId: uint64(this.AdminId()),
		},
		Id:    params.Id,
		Email: emails,
	})
	if err != nil || res.Code != 0 {
		if err == nil {
			err = fmt.Errorf(res.Msg)
		}
		this.ErrorPage(err)
		return
	}
	defer this.CreateLogInfo("修改 安全审计-主机 -授权 %v", res.Msg)

	this.Success()
}
