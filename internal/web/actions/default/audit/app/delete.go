package app

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_app"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	Id uint64
}) {
	res, err := audit_app.DelApp(&audit_app.DelAppReq{
		User: &request.UserReq{
			AdminUserId: uint64(this.AdminId()),
		},
		Id: params.Id,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	defer this.CreateLogInfo("删除安全审计-应用 %v", res.Msg)

	this.Success()
}
