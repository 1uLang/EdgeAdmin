package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
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
	Score       int
	MacCode     string
	ResultState int

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	this.Show()
	return
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &baseline.SearchReq{}
	req.UserName = params.UserName
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize
	req.MacCode = params.MacCode
	req.ResultState = params.ResultState
	if params.State > 0 {
		req.State = &params.State
	}
	list, err := baseline_server.List(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}
