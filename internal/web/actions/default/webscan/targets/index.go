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
	PageSize int
	PageNo   int

	Show int
}) {
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 20
	}
	list, err := targets_server.List(&targets.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, AdminUserId: uint64(this.AdminId())})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if lists, ok := list["targets"]; ok {
		this.Data["targets"] = lists
	}
	if params.Show == 1 {
		this.Success()
	}
	this.Show()
}
