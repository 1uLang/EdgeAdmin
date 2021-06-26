package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type CheckAction struct {
	actionutils.ParentAction
}

func (this *CheckAction) RunPost(params struct {
	MacCodes   []string `json:"macCodes"`
	TemplateId int      `json:"templateId"`

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("templateId", params.TemplateId).
		Require("请输入合规基线模板")

	if len(params.MacCodes) == 0 {
		this.ErrorPage(fmt.Errorf("请选择机器码"))
		return
	}

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &baseline.CheckReq{MacCodes: params.MacCodes, TemplateId: params.TemplateId}
	err = baseline_server.Check(req)

	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
