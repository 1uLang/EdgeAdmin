package reports

import (
	reports_server "github.com/1uLang/zhiannet-api/awvs/server/reports"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
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
	for _, reportId := range params.ReportIds {
		err = reports_server.Delete(reportId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	userResp, err := this.RPC().UserRPC().FindEnabledUser(this.AdminContext(), &pb.FindEnabledUserRequest{UserId: this.AdminId()})

	if err != nil {
		this.ErrorPage(err)
		return
	}
	user := userResp.User
	if user == nil {
		this.NotFound("user", this.AdminId())
		return
	}

	// 日志
	this.CreateLogInfo("WEB漏洞扫描 - 删除目标扫描报表:%v成功,用户名：%s", params.ReportIds, userResp.User.Username)
	this.Success()
}
