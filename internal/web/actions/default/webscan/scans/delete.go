package scans

import (
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ScanIds []string

	Must *actions.Must
}) {

	params.Must.
		Field("ScanIds", params.ScanIds).
		Require("请输入扫描id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for _, scanid := range params.ScanIds {
		err = scans_server.Delete(scanid)
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
	this.CreateLogInfo("WEB漏洞扫描 - 删除扫描任务目标:%v成功,用户名：%s", params.ScanIds, userResp.User.Username)
	this.Success()
}
