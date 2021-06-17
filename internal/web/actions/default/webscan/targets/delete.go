package targets

import (
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	targetId string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("targetId", params.targetId).
		Require("请输入目标id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	err = targets_server.Delete(params.targetId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
