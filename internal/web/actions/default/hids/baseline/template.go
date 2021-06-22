package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
)

type TemplateAction struct {
	actionutils.ParentAction
}

func (this *TemplateAction) Init() {
	this.FirstMenu("index")
}

func (this *TemplateAction) RunGet(params struct {
	PageNo   int
	PageSize int
	UserName string

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
	req := &baseline.TemplateSearchReq{}
	req.UserName = params.UserName
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize

	list, err := baseline_server.TemplateList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}

type TemplateDetailAction struct {
	actionutils.ParentAction
}

func (this *TemplateDetailAction) Init() {
	this.FirstMenu("index")
}

func (this *TemplateDetailAction) RunGet(params struct {
	PageNo     int
	PageSize   int
	UserName   string
	TemplateId string

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	this.Show()
	return
	//params.Must.
	//	Field("templateId", params.TemplateId).
	//	Require("请输入合规基线模板")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &baseline.TemplateDetailReq{}
	req.Req.UserName = params.UserName
	req.Req.PageNo = params.PageNo
	req.Req.PageSize = params.PageSize
	req.ID = params.TemplateId

	list, err := baseline_server.TemplateDetail(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = list
	this.Show()
}
