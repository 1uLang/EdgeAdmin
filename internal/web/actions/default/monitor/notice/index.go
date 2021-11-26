package notice

import (
	monitor_notice_server "github.com/1uLang/zhiannet-api/monitor/server/monitor_notice"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Message  string
	PageSize int
	PageNo   int
}) {
	list, total, err := monitor_notice_server.GetList(&monitor_notice_server.ListReq{Message: params.Message, PageNum: params.PageNo, PageSize: params.PageSize})
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["notices"] = list
	this.Data["total"] = total
	this.Data["message"] = params.Message
	this.Show()
	//this.Success()
}
