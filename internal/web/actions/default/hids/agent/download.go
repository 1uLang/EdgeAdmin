package agent

import (
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DownloadAction struct {
	actionutils.ParentAction
}

func (this *DownloadAction) Init() {
	this.FirstMenu("index")
}

func (this *DownloadAction) RunGet(params struct {
	UserName string
	osType   string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("username", params.UserName).
		Require("请输入用户名")

	params.Must.
		Field("osType", params.osType).
		Require("请选择主机操作系统")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	list, err := agent_server.Download(params.UserName, params.osType)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}
