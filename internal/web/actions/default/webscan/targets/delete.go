package targets

import (
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	TargetIds []string `json:"target_ids"`
}) {

	if len(params.TargetIds) == 0 {
		this.FailField("username", "请选择需要删除的目标")
		return
	}

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for _, targetId := range params.TargetIds {
		err = targets_server.Delete(targetId)
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
	this.CreateLogInfo("WEB漏洞扫描 - 删除任务目标:%v成功,用户名：%s", params.TargetIds, userResp.User.Username)

	this.Success()
}
