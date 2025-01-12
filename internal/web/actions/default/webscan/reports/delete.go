package reports

import (
	reports_server "github.com/1uLang/zhiannet-api/awvs/server/reports"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ReportIds []string

	Must *actions.Must
}) {

	params.Must.
		Field("ReportIds", params.ReportIds).
		Require("请输入报表id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	webscan_ids := []string{}
	hostscan_ids := []string{}
	for _, v := range params.ReportIds {
		if !strings.HasSuffix(v, "-host") { //去掉主机漏洞扫描
			webscan_ids = append(webscan_ids, v)
		} else {
			hostscan_ids = append(hostscan_ids, strings.TrimSuffix(v, "-host"))
		}
	}
	if len(hostscan_ids) > 0 {
		err = nessus_scans_server.DeleteReport(hostscan_ids)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	for _, reportId := range webscan_ids {
		err = reports_server.Delete(reportId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	// 日志
	this.CreateLogInfo("WEB漏洞扫描 - 删除目标扫描报表:%v成功", params.ReportIds)
	this.Success()
}
