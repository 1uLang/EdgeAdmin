package admins

import (
	"fmt"
	admin_users_model "github.com/1uLang/zhiannet-api/jumpserver/model/admin_users"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) checkAndNewServerRequest() (*jumpserver_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}
	username, _ := this.UserName()
	return fortcloud.NewServerRequest(username, "dengbao-"+username)
}
func (this *IndexAction) Init() {
	this.Nav("", "fortcloud", "index")
}

func (this *IndexAction) RunGet() {

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.Data["nodeErr"] = "堡垒机组件错误:" + err.Error()
		this.Show()
		return
	}

	list, err := req.AdminUser.List(&admin_users_model.ListReq{
		AdminUserId: uint64(this.AdminId()),
	})
	this.Data["adminUsers"] = list
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	Name     string
	Username string
	Password string
	Comment  string

	Must *actions.Must
}) {

	params.Must.
		Field("name", params.Name).
		Require("请输入名称")

	params.Must.
		Field("username", params.Username).
		Require("请输入用户名")

	params.Must.
		Field("password", params.Password).
		Require("请输入密码")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}

	args := &admin_users_model.CreateReq{}
	args.Name = params.Name
	args.UserName = params.Username
	args.Password = params.Password
	args.Comment = params.Comment
	args.AdminUserId = uint64(this.AdminId())
	_, err = req.AdminUser.Create(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 新增管理用户:[%v]成功", params.Name)
	this.Success()
}
