package admins

import (
	nc_req "github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	AdminId int64
}) {

	if params.AdminId == this.AdminId(){
		this.Fail("无权限")
		return
	}

	defer this.CreateLogInfo("删除系统用户 %d", params.AdminId)

	_, err := this.RPC().AdminRPC().DeleteAdmin(this.AdminContext(), &pb.DeleteAdminRequest{AdminId: params.AdminId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知更改
	err = configloaders.NotifyAdminModuleMappingChange()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 删除nc用户
	err = nc_req.DeleteUser(params.AdminId, 1)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
