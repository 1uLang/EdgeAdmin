package risk

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
}) {

	req := &risk.SearchReq{}
	req.UserName = "luobing"

	//系统漏洞数汇总
	risk, err := risk_server.SystemDistributed(req)
	if err != nil {
		this.ErrorPage(err)
	}
	//弱口令
	weak, err := risk_server.WeakList(req)
	if err != nil {
		this.ErrorPage(err)
	}
	//风险账号
	dangerAccount, err := risk_server.DangerAccountList(req)
	if err != nil {
		this.ErrorPage(err)
	}
	//缺陷配置
	configDefect, err := risk_server.ConfigDefectList(req)
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["risk"] = risk
	this.Data["weak"] = weak
	this.Data["dangerAccount"] = dangerAccount
	this.Data["configDefect"] = configDefect
	this.Show()
}
