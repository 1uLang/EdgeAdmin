package webscan

import (
	dashboard_server "github.com/1uLang/zhiannet-api/awvs/server/dashboard"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "webscan", "index")
}

func (this *IndexAction) RunGet() {
	err := InitAWVSServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	info, err := dashboard_server.MeState()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = info
	this.Show()
}
