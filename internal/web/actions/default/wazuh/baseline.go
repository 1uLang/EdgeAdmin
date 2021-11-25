// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type BaseLineAction struct {
	actionutils.ParentAction
}

func (this *BaseLineAction) Init() {
	this.Nav("", "", "baseline")
}

func (this *BaseLineAction) RunGet(params struct{}) {

	this.Data["baselines"] = []map[string]string{}
	this.Data["agents"] = []map[string]string{}
	defer this.Show()
	err := InitAPIServer()
	if err != nil {
		this.Data["errorMsg"] = err.Error()
		return
	}
	agent, err := server.AgentList(&agents.ListReq{
		AdminUserId: this.AdminId(),
	})
	if err != nil {
		this.Data["errorMsg"] = err.Error()
		return
	}
	if len(agent.AffectedItems) == 0 {
		this.Data["errorMsg"] = "请先添加资产"
		return
	}
	baselines := agents.SCAListResp{}
	for _, v := range agent.AffectedItems {
		list, err := server.BaselineList(agents.SCAListReq{
			Agent: v.ID,
		})
		if err != nil {
			this.Data["errorMsg"] = err.Error()
			return
		}
		for k := range list.AffectedItems {
			list.AffectedItems[k].AgentID = v.ID
			list.AffectedItems[k].AgentIP = v.IP
			list.AffectedItems[k].AgentName = v.Name
		}
		baselines.AffectedItems = append(baselines.AffectedItems, list.AffectedItems...)
	}
	this.Data["baselines"] = baselines.AffectedItems

	this.Data["agents"] = agent.AffectedItems

}

type BaseLineDetailsAction struct {
	actionutils.ParentAction
}

func (this *BaseLineDetailsAction) Init() {
	this.Nav("", "", "details")
}

func (this *BaseLineDetailsAction) RunGet(params struct {
	Agent  string
	Policy string
}) {
	list, err := server.BaselineDetailsList(agents.SCADetailsListReq{
		Agent:  params.Agent,
		Policy: params.Policy,
		Limit:  1,
		Offset: 0,
		Result: "failed",
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	page := this.NewPage(list.TotalAffectedItems)
	this.Data["page"] = page.AsHTML()

	list, err = server.BaselineDetailsList(agents.SCADetailsListReq{
		Agent:  params.Agent,
		Policy: params.Policy,
		Limit:  page.Size,
		Offset: page.Offset,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["details"] = list.AffectedItems
	this.Show()
}
