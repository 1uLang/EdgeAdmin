package webscan

import (
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/awvs/server"
	dashboard_server "github.com/1uLang/zhiannet-api/awvs/server/dashboard"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Keyword string
}) {
	err := server.SetUrl("https://scan-web.zhiannet.com/")
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//初始化 awvs 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		XAuth: "1986ad8c0a5b3df4d7028d5f3c06e936c429ffb1149e2491b84fe51cc63a6b26a",
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	info, err := dashboard_server.MeState()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = info
	this.Show()
}
