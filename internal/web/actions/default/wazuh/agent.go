// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type AgentsAction struct {
	actionutils.ParentAction
}

func (this *AgentsAction) Init() {
	this.Nav("", "", "agents")
}

func (this *AgentsAction) RunGet(params struct{}) {
	defer this.Show()
	this.Data["agents"] = []int64{}
	err := InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	list, err := server.AgentList(&agents.ListReq{
		AdminUserId: this.AdminId(),
	})
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	this.Data["agents"] = list.AffectedItems

}

func (this *AgentsAction) RunPost(params struct {
	Agent string
}) {

	err := InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	err = server.AgentDelete([]string{params.Agent})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.CreateLogInfo("主机防护 - 删除资产:[%v]成功", params.Agent)
	this.Success()
}
