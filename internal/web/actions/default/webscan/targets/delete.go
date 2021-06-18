package targets

import (
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	TargetIds []string `json:"target_ids"`
}) {

	if len(params.TargetIds) == 0 {
		this.FailField("username", "请选择需要删除的目标")
		return
	}

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for _, targetId := range params.TargetIds {
		err = targets_server.Delete(targetId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	this.Success()
}
