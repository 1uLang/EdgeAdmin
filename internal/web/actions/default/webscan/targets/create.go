package targets

import (
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
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
	Address string
	Desc    string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("address", params.Address).
		Require("请输入目标地址")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &targets.AddReq{Address: params.Address, AdminUserId: uint64(this.AdminId())}
	req.Description = params.Desc
	//req.AdminUserId = uint64(this.AdminId())

	_, err = targets_server.Add(req)
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
	this.CreateLogInfo("WEB漏洞扫描 - 创建任务目标:[%v]成功", params.Address)
	this.Success()
}
