// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type AttckAction struct {
	actionutils.ParentAction
}

func (this *AttckAction) Init() {
	this.Nav("", "", "virus")
}

func (this *AttckAction) RunGet(params struct {
	Agent string
}) {

	this.Data["attcks"] = []map[string]string{}
	this.Data["agents"] = []map[string]string{}

	this.Data["agent"] = params.Agent

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
	if params.Agent == "" {
		params.Agent = agent.AffectedItems[0].ID
		//params.Agent = "007"
	}
	list, err := server.ATTCKESList(agents.ESListReq{
		Agent:  params.Agent,
		Limit:  1,
		Offset: 0,
		//Start:  time.Now().AddDate(0, 0, -1).Unix(),
		//End:    time.Now().Unix(),
		//Start: 1630982235, End: 1631068635,
	})
	if err != nil {
		this.Data["errorMsg"] = err.Error()
		return
	}

	page := this.NewPage(int64(list.Total))
	this.Data["page"] = page.AsHTML()

	list, err = server.ATTCKESList(agents.ESListReq{
		Agent:  params.Agent,
		Limit:  int(page.Size),
		Offset: int(page.Offset),
		//Start:  time.Now().AddDate(0, 0, -1).Unix(),
		//End:    time.Now().Unix(),
		//Start: 1630982235, End: 1631068635,
	})
	if err != nil {
		this.Data["errorMsg"] = err.Error()
		return
	}
	this.Data["attcks"] = list.Hits

	this.Data["agents"] = agent.AffectedItems

	this.Data["agent"] = params.Agent

}
