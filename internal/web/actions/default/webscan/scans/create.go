package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}
func (this *CreateAction) RunPost(params struct {
	TargetIds []string
}) {

	if len(params.TargetIds) == 0 {
		this.FailField("username", "请选择需要扫描的目标")
		return
	}

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &scans.AddReq{ProfileId: "11111111-1111-1111-1111-111111111111"}
	for _, targetId := range params.TargetIds {
		req.TargetId = targetId
		err = scans_server.Add(req)
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
	this.CreateLogInfo("WEB漏洞扫描 - 扫描任务目标:%v成功,用户名：%s", params.TargetIds, userResp.User.Username)
	this.Success()
}
