package monitor

import (
	monitor_list_server "github.com/1uLang/zhiannet-api/monitor/server/monitor_list"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	NShowState int
	PageSize   int
	PageNo     int
}) {
	//默认端口监控
	if params.NShowState == 0 {
		params.NShowState = 1
	}
	list, total, err := monitor_list_server.GetList(&monitor_list_server.ListReq{MonitorType: params.NShowState, PageSize: params.PageSize, PageNum: params.PageNo})
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["monitors"] = list
	this.Data["total"] = total
	this.Data["nShowState"] = params.NShowState
	this.Show()

}
