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

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}
func (this *CreateAction) RunPost(params struct {
	TargetId string `json:"target_id"`

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("target_id", params.TargetId).
		Require("请输入目标id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &scans.AddReq{TargetId: params.TargetId, ProfileId: "11111111-1111-1111-1111-111111111111"}

	err = scans_server.Add(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
