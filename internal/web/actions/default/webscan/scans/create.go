package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) PostGet(params struct {
	TargetId  string
	ProfileId string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("target_id", params.TargetId).
		Require("请输入目标id")

	params.Must.
		Field("profile_id", params.ProfileId).
		Require("请输入扫描类型")

	err := webscan.InitAWVSServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &scans.AddReq{TargetId: params.TargetId, ProfileId: params.ProfileId}

	err = scans_server.Add(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
