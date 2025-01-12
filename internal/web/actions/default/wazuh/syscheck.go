// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type SysCheckAction struct {
	actionutils.ParentAction
}

func (this *SysCheckAction) Init() {
	this.Nav("", "", "virus")
}

func (this *SysCheckAction) RunGet(params struct {
	Agent string
	Event string
	Path  string
	File  string //查询文件路径
}) {
	this.Data["event"] = params.Event
	this.Data["filePath"] = params.File
	this.Data["syschecks"] = []map[string]string{}
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
	if params.Event == "" { //文件列表
		list, err := server.SysCheckList(agents.SysCheckListReq{
			Agent:  params.Agent,
			Limit:  1,
			Offset: 0,
			File:   params.File,
			//Start: 1630982235, End: 1631068635,
		})
		if err != nil {
			this.Data["errorMsg"] = err.Error()
			return
		}

		page := this.NewPage(list.TotalAffectedItems)
		this.Data["page"] = page.AsHTML()

		list, err = server.SysCheckList(agents.SysCheckListReq{
			Agent:  params.Agent,
			Limit:  page.Size,
			Offset: page.Offset,
			File:   params.File,
		})
		if err != nil {
			this.Data["errorMsg"] = err.Error()
			return
		}
		this.Data["syschecks"] = list.AffectedItems

		this.Data["agents"] = agent.AffectedItems

		this.Data["agent"] = params.Agent
		return
	}

	list, err := server.SysCheckESList(agents.ESListReq{
		Agent:  params.Agent,
		Path:   params.Path,
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

	list, err = server.SysCheckESList(agents.ESListReq{
		Agent:  params.Agent,
		Path:   params.Path,
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
	this.Data["syschecks"] = list.Hits

	this.Data["agents"] = agent.AffectedItems

	this.Data["agent"] = params.Agent

}
