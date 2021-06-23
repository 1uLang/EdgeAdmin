package agent

import (
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
)

type InstallAction struct {
	actionutils.ParentAction
}

func (this *InstallAction) Init() {
	this.FirstMenu("index")
}

func (this *InstallAction) RunGet(params struct {
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	linux, err := agent_server.Install("luobing", "Linux")
	if err != nil {
		this.ErrorPage(err)
		return
	}

	windows, err := agent_server.Install("luobing", "Windows")
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["linux"] = linux
	this.Data["windows"] = windows
	this.Show()
}
