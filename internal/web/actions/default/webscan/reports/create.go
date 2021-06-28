package reports

import (
	"github.com/1uLang/zhiannet-api/awvs/model/reports"
	reports_server "github.com/1uLang/zhiannet-api/awvs/server/reports"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}
func (this *CreateAction) RunPost(params struct {
	Ids    []string
	TarIds []string
	Must   *actions.Must
}) {

	params.Must.
		Field("id_list", params.Ids).
		Require("请选择指定生成报表的扫描目标")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &reports.CreateResp{
		Source: struct {
			IDS  []string `json:"id_list"`
			Type string   `json:"list_type"`
		}{IDS: params.Ids, Type: "scans"},
		TemplateId:  "11111111-1111-1111-1111-111111111112", //快速
		AdminUserId: uint64(this.AdminId()),
	}
	_, err = reports_server.Create(req)
	if err != nil {
		this.ErrorPage(err)
		return
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
	this.CreateLogInfo("WEB漏洞扫描 - 生成目标扫描报表:%v成功,用户名：%s", params.Ids, userResp.User.Username)
	this.Success()
}
