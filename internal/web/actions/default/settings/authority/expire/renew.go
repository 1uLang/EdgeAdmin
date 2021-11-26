package expire

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/index"
	"github.com/iwind/TeaGo/actions"
)

type RenewAction struct {
	actionutils.ParentAction
}

func (this *RenewAction) Init() {
	this.Nav("", "", "index")
}

func (this *RenewAction) RunGet(params struct{}) {

	var err error
	this.Data["systemCode"], _, err = index.CheckExpire()
	if err != nil {
		this.ErrorPage(err)
	}

	this.Show()
}

func (this *RenewAction) RunPost(params struct {
	Secret string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.Field("secret", params.Secret).Require("请输入续订密钥")

	err := index.Renew(params.Secret)
	if err != nil {
		// 日志
		this.CreateLogInfo("系统续订失败:%v", err)
		this.FailField("secret", err.Error())
	}
	this.Success()
}
