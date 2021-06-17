package targets

import (
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
)

//任务目标
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
	err := webscan.InitAPIServer()
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
	list, err := targets_server.List(&targets.ListReq{Limit: params.pageSize, C: params.pageNo * params.pageSize})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["targets"] = list["targets"]
	this.Show()
}
