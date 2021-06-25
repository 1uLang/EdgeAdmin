package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	examine_server "github.com/1uLang/zhiannet-api/hids/server/examine"
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
	PageNo       int
	PageSize     int
	UserName     string
	State        int
	Score        int
	Type         int
	StartTime    string   //体检开始时间
	EndTime      string   //体检结束时间
	ExamineItems []string //体检项目集合

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
	}
	if params.State == 0 {
		params.State = -1
	}
	if params.Score == 0 {
		params.Score = -1
	}
	if params.Type == 0 {
		params.Type = -1
	}
	req := &examine.SearchReq{}
	req.UserName = "luobing"
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize

	req.State = params.State
	req.Score = params.Score
	req.Type = params.Type

	req.StartTime = params.StartTime
	req.EndTime = params.EndTime
	req.ExamineItems = params.ExamineItems

	list, err := examine_server.List(req)
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["data"] = list
	this.Data["state"] = params.State
	this.Data["type"] = params.Type
	this.Data["score"] = params.Score

	this.Show()
}
