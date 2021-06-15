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
	pageSize int
	pageNo   int
}) {

	err := webscan.InitAWVSServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.pageNo < 0 {
		params.pageNo = 0
	}
	if params.pageSize < 0 {
		params.pageSize = 20
	}
	list, err := scans_server.List(&scans.ListReq{Limit: params.pageSize, C: params.pageNo * params.pageSize})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["scans"] = list["scans"]
	this.Success()
}
