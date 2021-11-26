package db

import (
	"github.com/1uLang/zhiannet-api/common/model/audit_assets_relation"
	"github.com/1uLang/zhiannet-api/common/server/audit_assets_relation_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_admins_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type AuthorizeAction struct {
	actionutils.ParentAction
}

type UserList struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	IsOn bool   `json:"is_on"`
	My   bool   `json:"my"`
}

func (this *AuthorizeAction) RunGet(params struct {
	Id   uint64
	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择数据库")

	//获取管理端用户
	list, _, err := edge_admins_server.GetList()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	//获取授权中的列表
	authUse, _, err := audit_assets_relation_server.GetList(&audit_assets_relation.ListReq{
		//AdminUserId: uint64(this.AdminId()),
		AssetsId:   []uint64{params.Id},
		AssetsType: 0,
		PageSize:   999,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	authMap := map[uint64]uint64{}
	if len(authUse) > 0 {
		for _, v := range authUse {
			authMap[v.AdminUserId] = v.AdminUserId
		}
	}
	allUsers, authUsers := make([]UserList, 0), make([]UserList, 0)
	if len(list) > 0 {
		for _, v := range list {
			var isOn bool
			if _, ok := authMap[v.Id]; ok {
				isOn = true
			}
			if isOn {
				if v.Id == uint64(this.AdminId()) {
					allUsers = append(allUsers, UserList{
						Id:   v.Id,
						Name: v.Username,
						IsOn: isOn,
						My:   v.Id == uint64(this.AdminId()),
					})
				} else {
					authUsers = append(authUsers, UserList{
						Id:   v.Id,
						Name: v.Username,
						IsOn: isOn,
						My:   v.Id == uint64(this.AdminId()),
					})
				}

			} else {
				allUsers = append(allUsers, UserList{
					Id:   v.Id,
					Name: v.Username,
					IsOn: isOn,
					My:   v.Id == uint64(this.AdminId()),
				})
			}

		}
	}

	//
	//
	//
	//list, err := audit_db.GetAuthEmail(&server.AuthReq{
	//	Id: params.Id,
	//	User: &request.UserReq{
	//		AdminUserId: uint64(this.AdminId()),
	//	},
	//})
	////var email string
	//allUsers, authUsers := make([]UserList, 0), make([]UserList, 0)
	//if err != nil || list == nil {
	//
	//} else {
	//	for _, v := range list.Data.UserList {
	//		if v.IsOn {
	//
	//			if v.My {
	//				allUsers = append(allUsers, UserList{
	//					Id:   v.Id,
	//					Name: v.Name,
	//					My:   v.My,
	//					IsOn: v.IsOn,
	//				})
	//			} else {
	//				authUsers = append(authUsers, UserList{
	//					Id:   v.Id,
	//					Name: v.Name,
	//					My:   v.My,
	//					IsOn: v.IsOn,
	//				})
	//			}
	//		} else {
	//			allUsers = append(allUsers, UserList{
	//				Id:   v.Id,
	//				Name: v.Name,
	//				My:   v.My,
	//				IsOn: v.IsOn,
	//			})
	//		}
	//
	//	}
	//}
	//email = strings.TrimSpace(email)
	this.Data["allUsers"] = allUsers
	this.Data["authUsers"] = authUsers
	this.Success()
}
func (this *AuthorizeAction) RunPost(params struct {
	Id    uint64
	Users []uint64
	Must  *actions.Must
}) {
	if len(params.Users) == 0 {
		params.Users = []uint64{uint64(this.AdminId())}
	}

	params.Must.
		Field("id", params.Id).
		Require("请选择数据库")

	//res, err := audit_db.AuthDb(&server.AuthReq{
	//	User: &request.UserReq{
	//		AdminUserId: uint64(this.AdminId()),
	//	},
	//	Id:  params.Id,
	//	Ids: params.Users,
	//	//Email: emails,
	//})
	//if err != nil || res.Code != 0 {
	//	if err == nil {
	//		err = fmt.Errorf(res.Msg)
	//	}
	//	this.ErrorPage(err)
	//	return
	//}
	users := append([]uint64{uint64(this.AdminId())}, params.Users...)
	res := audit_assets_relation_server.Reset(&audit_assets_relation.AddReq{
		AdminUserId: users,
		AssetsId:    params.Id,
		AssetsType:  0,
	})
	defer this.CreateLogInfo("修改 安全审计-数据库 -授权 %v", res)
	if res != nil {
		this.ErrorPage(res)
		return
	}
	this.Success()
}
