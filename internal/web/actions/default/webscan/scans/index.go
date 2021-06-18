package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageSize int
	PageNo   int
}) {

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.PageNo <= 0 {
		params.PageNo = 0
	}
	if params.PageSize <= 0 {
		params.PageSize = 20
	}
	list, err := scans_server.List(&scans.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["scans"] = list["scans"]
	this.Show()
}
