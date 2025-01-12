package report

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_from"
	"github.com/1uLang/zhiannet-api/common/model/audit_assets_relation"
	"github.com/1uLang/zhiannet-api/common/server/audit_assets_relation_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"time"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct {
	Id uint64
}) {
	if params.Id > 0 {
		info, _ := audit_from.GetFrom(&audit_from.GetFromReq{
			User: &request.UserReq{
				AdminUserId: uint64(this.AdminId()),
			},
			Id: params.Id,
		})
		this.Data["id"] = info.Data.Info.ID
		this.Data["name"] = info.Data.Info.Name
		this.Data["email"] = info.Data.Info.Email
		this.Data["format"] = info.Data.Info.Format
		this.Data["assets_id"] = info.Data.Info.AssetsID
		this.Data["assets_type"] = info.Data.Info.AssetsType
		this.Data["cycle"] = info.Data.Info.Cycle
		this.Data["cycle_day"] = info.Data.Info.CycleDay
		this.Data["send_time"] = info.Data.Info.SendTime
	} else {

		this.Data["id"] = 0
		this.Data["name"] = ""
		this.Data["email"] = ""
		this.Data["format"] = 1
		this.Data["assets_id"] = -1
		this.Data["assets_type"] = 1
		this.Data["cycle"] = 1
		this.Data["cycle_day"] = 1
		this.Data["send_time"] = "00:00"
	}

	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name       string
	Email      string
	Format     int
	AssetsType int
	AssetsId   int64
	Cycle      int
	CycleDay   int
	SendTime   string
	Id         uint64

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入名称")
	params.Must.
		Field("email", params.Email).
		Require("请输入邮箱")
	params.Must.
		Field("assetsid", params.AssetsId).
		Gt(0, "请选择资产")
	if params.Id == 0 {
		res, err := audit_from.AddFrom(&audit_from.FromReq{
			User: &request.UserReq{
				AdminUserId: uint64(this.AdminId()),
			},
			Name:       params.Name,
			Email:      params.Email,
			Format:     params.Format,
			AssetsType: params.AssetsType,
			AssetsId:   uint64(params.AssetsId),
			Cycle:      params.Cycle,
			CycleDay:   params.CycleDay,
			SendTime:   params.SendTime,
		})

		if err != nil {
			this.ErrorPage(err)
			return
		}
		if res != nil && res.Code != 0 {
			this.Fail(res.Msg)
			return
		}
		//关联账号审计ID
		if res != nil {
			_, err = audit_assets_relation_server.Add(&audit_assets_relation.AuditAssetsRelation{
				AdminUserId: uint64(this.AdminId()),
				AssetsId:    res.Data.Id,
				AssetsType:  4, //数据报表
				CreateTime:  int(time.Now().Unix()),
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}
		defer this.CreateLogInfo("编辑安全审计-报表 %v", res.Msg)
	} else {
		res, err := audit_from.EditFrom(&audit_from.FromReq{
			User: &request.UserReq{
				AdminUserId: uint64(this.AdminId()),
			},
			Name:       params.Name,
			Email:      params.Email,
			Format:     params.Format,
			AssetsType: params.AssetsType,
			AssetsId:   uint64(params.AssetsId),
			Cycle:      params.Cycle,
			CycleDay:   params.CycleDay,
			SendTime:   params.SendTime,
			Id:         params.Id,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if res != nil && res.Code != 0 {
			this.Fail(res.Msg)
			return
		}
		defer this.CreateLogInfo("创建安全审计-报表 %v", res.Msg)
	}
	this.Success()
}
