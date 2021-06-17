package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) RunGet(params struct {
	MacCode  string
	PageNo   int
	PageSize int

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")

	if err := hids.InitAPIServer(); err != nil {
		this.ErrorPage(err)
		return
	}

	req := &baseline.DetailReq{}
	req.MacCode = params.MacCode
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize
	info, err := baseline_server.Detail(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = info

	this.Show()
}
