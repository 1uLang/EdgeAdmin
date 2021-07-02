package risk

import (
	"fmt"
	risk_model "github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
}) {

	defer this.Show()

	var risk, weak, dangerAccount, configDefect risk_model.SystemDistributedResp
	this.Data["risk"] = risk
	this.Data["weak"] = weak
	this.Data["dangerAccount"] = dangerAccount
	this.Data["configDefect"] = configDefect

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &risk_model.SearchReq{}

	req.UserName, err = this.UserName()
	if err != nil {
		this.Data["errorMessage"] = fmt.Errorf("获取用户信息失败：%v", err)
		return
	}

	//系统漏洞数汇总
	risk, err = risk_server.SystemDistributed(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Errorf("获取系统漏洞信息失败：%v", err)
		return
	}
	//弱口令
	weak, err = risk_server.WeakList(req)
	if err != nil {

		this.Data["errorMessage"] = fmt.Errorf("获取弱口令信息失败：%v", err)
		return
	}
	//风险账号
	dangerAccount, err = risk_server.DangerAccountList(req)
	if err != nil {

		this.Data["errorMessage"] = fmt.Errorf("获取风险账号信息失败：%v", err)
		return
	}
	//缺陷配置
	configDefect, err = risk_server.ConfigDefectList(req)
	if err != nil {

		this.Data["errorMessage"] = fmt.Errorf("获取缺陷配置信息失败：%v", err)
		return
	}
	this.Data["risk"] = risk
	this.Data["weak"] = weak
	this.Data["dangerAccount"] = dangerAccount
	this.Data["configDefect"] = configDefect

}
