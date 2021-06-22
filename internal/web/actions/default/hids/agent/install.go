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
	UserName string
	osType   string

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	this.Show()
	return
	//params.Must.
	//	Field("username", params.UserName).
	//	Require("请输入用户名")
	//
	//params.Must.
	//	Field("osType", params.osType).
	//	Require("请选择主机操作系统")
	//
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	list, err := agent_server.Install(params.UserName, params.osType)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}
