package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	examine_server "github.com/1uLang/zhiannet-api/hids/server/examine"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type ScanAction struct {
	actionutils.ParentAction
}

func (this *ScanAction) RunPost(params struct {
	Opt       string
	MacCode   []string
	ScanItems []string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("opt", params.Opt).
		Require("请输入操作方式")

	if params.Opt != "now" && params.Opt != "cancel" {
		this.ErrorPage(fmt.Errorf("操作方式参数错误"))
		return
	}

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.Opt == "now" {
		err = examine_server.ScanServerNow(&examine.ScanReq{MacCode: params.MacCode, ScanItems: params.ScanItems})
	} else {
		err = examine_server.ScanServerCancel(params.MacCode)
	}

	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
