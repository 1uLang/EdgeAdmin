package agent

import (
	"github.com/1uLang/zhiannet-api/hids/model/agent"
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageNo        int
	PageSize      int
	UserName      string
	ServerIp      string
	ServerLocalIp string

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &agent.SearchReq{}
	//req.UserName = params.UserName
	req.UserName = "luobing"
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize
	req.ServerIp = params.ServerIp
	req.ServerLocalIp = params.ServerLocalIp

	list, err := agent_server.List(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["agents"] = list.List
	this.Show()
}
