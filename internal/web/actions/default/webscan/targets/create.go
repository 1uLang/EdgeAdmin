package targets

import (
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/logs"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) PostGet(params struct {
	Address string
	Desc    string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("address", params.Address).
		Require("请输入目标地址")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &targets.AddReq{Address: params.Address}
	req.Description = params.Desc

	targetId, err := targets_server.Add(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	logs.Infof("新建目标成功 ：%v", targetId)
	this.Success()
}
