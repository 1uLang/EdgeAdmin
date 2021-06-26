package scans

import (
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

//

type StatisticsAction struct {
	actionutils.ParentAction
}

func (this *StatisticsAction) Init() {
	this.FirstMenu("index")
}

func (this *StatisticsAction) RunGet(params struct {
	ScanId        string
	ScanSessionId string
	TargetId      string

	Must *actions.Must
}) {

	params.Must.
		Field("ScanId", params.ScanId).
		Require("请输入扫描id")

	params.Must.
		Field("ScanSessionId", params.ScanSessionId).
		Require("请输入扫描会话id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	info, err := scans_server.Statistics(params.ScanId, params.ScanSessionId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["statistics"] = info
	this.Success()
}
