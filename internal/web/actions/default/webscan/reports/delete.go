package reports

import (
	reports_server "github.com/1uLang/zhiannet-api/awvs/server/reports"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ReportId string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("report_id", params.ReportId).
		Require("请输入扫描id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	err = reports_server.Delete(params.ReportId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
