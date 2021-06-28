package webscan

import (
	dashboard_server "github.com/1uLang/zhiannet-api/awvs/server/dashboard"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "webscan", "index")
}

func (this *IndexAction) RunGet() {
	err := InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	info, err := dashboard_server.MeState()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = info

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
	this.CreateLogInfo("WEB漏洞扫描请求成功,用户名：%s", userResp.User.Username)
	this.Show()
}
