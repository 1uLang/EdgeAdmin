package users

import (
	"github.com/1uLang/zhiannet-api/nextcloud"
	nc_req "github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	UserId int64
}) {
	defer this.CreateLogInfo("删除用户 %d", params.UserId)

	// TODO 检查用户是否有未完成的业务

	_, err := this.RPC().UserRPC().DeleteUser(this.AdminContext(), &pb.DeleteUserRequest{UserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if nextcloud.UseDatabackup {

		err = nc_req.DeleteUser(params.UserId, 0)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	this.Success()
}
