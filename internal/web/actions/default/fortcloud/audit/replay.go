package audit

import (
	"fmt"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type ReplayAction struct {
	actionutils.ParentAction
}

func (this *ReplayAction) checkAndNewServerRequest() (*jumpserver_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}
	username, _ := this.UserName()
	return fortcloud.NewServerRequest(username, "dengbao-"+username)
	//return fortcloud.NewServerRequest("admin", "21ops.com")
}
func (this *ReplayAction) RunPost(params struct {
	Id   string
	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择需要回放的会话")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	url, token, err := req.Session.Replay(params.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["url"] =  fortcloud.ServerUrl+url
	this.Data["token"] = token
	this.Success()
}