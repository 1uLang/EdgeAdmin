package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type TemplateAction struct {
	actionutils.ParentAction
}

func (this *TemplateAction) Init() {
	this.FirstMenu("index")
}

func (this *TemplateAction) RunGet(params struct {
	MacCode string
	Os      string
	Must    *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.Field("macCode", params.MacCode).Require("请输入机器码")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &baseline.TemplateSearchReq{}
	req.PageNo = 1
	req.PageSize = 100

	list, err := baseline_server.TemplateList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	ls := make([]map[string]interface{}, 0, len(list.List))
	check := func(all, find string) bool {
		if len(all) < len(find) {
			all, find = find, all
		}
		return strings.Contains(strings.ToUpper(all), strings.ToUpper(find))
	}
	//todo:列出当前机器码对应主机系统的模板
	for _, v := range list.List {
		if v["type"].(float64) == 4 && check(params.Os, "win") {
			ls = append(ls, v)
		} else if v["type"].(float64) == 3 && !check(params.Os, "win") {
			ls = append(ls, v)
		}
	}
	this.Data["templates"] = ls

	this.Data["macCode"] = params.MacCode
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
