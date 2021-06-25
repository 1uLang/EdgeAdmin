package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
	"github.com/1uLang/zhiannet-api/hids/server/server"
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
	PageNo      int
	PageSize    int
	UserName    string
	State       int
	ResultState int

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &baseline.SearchReq{}
	req.UserName = "luobing"
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize
	req.ResultState = params.ResultState

	state := 0
	if params.State > 0 {
		state = params.State - 1
		req.State = &state
	}
	list, err := baseline_server.List(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for k, v := range list.List {
		os, err := server.Info(v["serverIp"].(string))
		if err != nil {
			this.ErrorPage(err)
		}
		list.List[k]["os"] = os
	}
	this.Data["baselines"] = list.List
	this.Data["State"] = params.State
	this.Data["ResultState"] = params.ResultState

	this.Show()
}
