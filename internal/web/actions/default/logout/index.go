package logout

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/rpc"
	"github.com/TeaOSLab/EdgeAdmin/internal/utils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction actions.Action

// 退出登录
func (this *IndexAction) Run(params struct {
	Auth *helpers.UserShouldAuth
}) {
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		this.Fail("服务器出了点小问题：" + err.Error())
	}
	userResp, err := rpcClient.UserRPC().FindEnabledUser(rpcClient.Context(int64(params.Auth.AdminId())), &pb.FindEnabledUserRequest{UserId: int64(params.Auth.AdminId())})
	if err != nil {
		this.Fail("获取用户信息失败")
		return
	}
	// 记录日志
	err = dao.SharedLogDAO.CreateAdminLog(rpcClient.Context(int64(params.Auth.AdminId())), oplogs.LevelInfo, this.Request.URL.Path, "退出系统，用户名："+userResp.User.Username, this.RequestRemoteIP())
	if err != nil {
		utils.PrintError(err)
	}

	params.Auth.Logout()
	this.RedirectURL("/")
}
