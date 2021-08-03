package report

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_from"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	Id uint64
}) {
	res, err := audit_from.DelFrom(&audit_from.DelFromReq{
		User: &request.UserReq{
			AdminUserId: uint64(this.AdminId()),
		},
		Id: params.Id,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	defer this.CreateLogInfo("删除安全审计-报表 %v", res.Msg)

	this.Success()
}
