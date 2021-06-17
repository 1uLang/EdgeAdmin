package examine

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type ScanAction struct {
	actionutils.ParentAction
}

func (this *ScanAction) RunPost(params struct {
	ProviderId int64
	Opt        string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("opt", params.Opt).
		Require("请输入操作方式")

	if params.Opt != "now" && params.Opt != "cancel" {
		this.ErrorText("操作方式参数错误")
		return
	}

	this.Success()
}
