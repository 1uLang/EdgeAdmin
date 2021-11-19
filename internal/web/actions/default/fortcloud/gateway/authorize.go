package gateway

import (
	"fmt"
	gateway_model "github.com/1uLang/zhiannet-api/next-terminal/model/access_gateway"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/fortcloud"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
)

type AuthorizeAction struct {
	actionutils.ParentAction
}

func (this *AuthorizeAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}

func (this *AuthorizeAction) RunGet(params struct {
	Id   string
	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择网关")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	list, err := req.GateWay.AuthorizeUserList(params.Id)
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	contain := map[uint64]bool{}
	for _, v := range list[1:] { //从一开始 去掉创建者
		contain[v] = true
	}
	countResp, err := this.RPC().AdminRPC().CountAllEnabledAdmins(this.AdminContext(), &pb.CountAllEnabledAdminsRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	adminsResp, err := this.RPC().AdminRPC().ListEnabledAdmins(this.AdminContext(), &pb.ListEnabledAdminsRequest{
		Offset: 0,
		Size:   countResp.Count,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	allUsers := make([]map[string]interface{}, 0)
	authUsers := make([]map[string]interface{}, 0)
	for _, v := range adminsResp.Admins {
		if _, isExist := contain[fmt.Sprintf("%v", v.Id)]; isExist {
			authUsers = append(authUsers, map[string]interface{}{
				"name": v.Username,
				"id":   v.Id,
			})
		} else {
			allUsers = append(allUsers, map[string]interface{}{
				"name": v.Username,
				"id":   v.Id,
				"my":   v.Id == this.AdminId(),
			})
		}
	}
	this.Data["allUsers"] = allUsers
	this.Data["authUsers"] = authUsers
	this.Success()
}
func (this *AuthorizeAction) RunPost(params struct {
	Id    string
	Users []uint64
	Must  *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择资产")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	args := &gateway_model.AuthorizeReq{}
	args.Id = params.Id
	args.AdminUserIds = params.Users
	args.AdminUserId = uint64(this.AdminId())
	err = req.GateWay.Authorize(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 修改网关授权:[%v]成功", params.Id)
	this.Success()
}
