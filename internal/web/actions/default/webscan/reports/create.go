package reports

import (
	"github.com/1uLang/zhiannet-api/awvs/model/reports"
	reports_server "github.com/1uLang/zhiannet-api/awvs/server/reports"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/logs"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) PostGet(params struct {
	Ids []string `json:"id_list"`

	TemplateId string `json:"template_id"`

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("id_list", params.Ids).
		Require("请选择指定生成报表的扫描目标")

	params.Must.
		Field("template_id", params.TemplateId).
		Require("请选择生成的报表模板")

	err := webscan.InitAWVSServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &reports.CreateResp{
		Source: struct {
			IDS  []string `json:"id_list"`
			Type string   `json:"list_type"`
		}{IDS: params.Ids, Type: "scans"},
		TemplateId: params.TemplateId,
	}
	info, err := reports_server.Create(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	logs.Infof("生成目标扫描报表成功 ：%v", info[""])
	this.Success()
}
