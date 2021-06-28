package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	examine_server "github.com/1uLang/zhiannet-api/hids/server/examine"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type ScanAction struct {
	actionutils.ParentAction
}

func (this *ScanAction) RunPost(params struct {
	Opt       string
	MacCode   []string
	ScanItems string

	VirusPath    string
	WebShellPath string
	Must         *actions.Must
	//CSRF         *actionutils.CSRF
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
		req := &examine.ScanReq{MacCode: params.MacCode}
		//去掉 ','
		params.ScanItems = strings.TrimPrefix(params.ScanItems, ",")
		params.ScanItems = strings.TrimSuffix(params.ScanItems, ",")
		if len(params.ScanItems) > 0 {
			req.ScanItems = strings.Split(params.ScanItems, ",")
		}
		req.ScanConfig.VirusPath = params.VirusPath
		req.ScanConfig.WebShellPath = params.WebShellPath
		err = examine_server.ScanServerNow(req)
	} else {
		err = examine_server.ScanServerCancel(params.MacCode)
	}

	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
