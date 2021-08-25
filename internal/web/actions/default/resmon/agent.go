package resmon

import (
	"github.com/1uLang/zhiannet-api/resmon/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type AgentAction struct {
	actionutils.ParentAction
}

func (this *AgentAction) Init() {
	this.Nav("", "", "")
}

func (this *AgentAction) RunGet(params struct {
	Id   string
	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("请输入id")

	downInfo := server.GetNodeURL(params.Id)

	this.Data["dwon_info"] = downInfo
	this.Success()
}
