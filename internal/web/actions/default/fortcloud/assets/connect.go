package assets

import (
	"fmt"
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type ConnectAction struct {
	actionutils.ParentAction
}

func (this *ConnectAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}

func (this *ConnectAction) RunPost(params struct {
	Id   string
	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择资产")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	args := &asset_model.ConnectReq{}
	args.Id = params.Id
	url, err := req.Assets.Connect(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["url"] = fortcloud.ServerUrl + url
	// 日志
	this.CreateLogInfo("堡垒机 - 连接资产:[%v]成功", params.Id)
	this.Success()
}
