package scans

import (
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	scanId string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("scanId", params.scanId).
		Require("请输入扫描id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	err = scans_server.Delete(params.scanId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
