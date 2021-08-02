package db

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_db"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "audit", "index")
}

func (this *IndexAction) RunGet(params struct {
	PageNum  int
	PageSize int
	Type     string
	Ip       string
	Name     string
	Status   string
	Json     bool

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {

	list, _ := audit_db.GetAuditBdList(&audit_db.ReqSearch{
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
		Type:     params.Type,
		Ip:       params.Ip,
		Name:     params.Name,
		Status:   params.Status,
		User: &request.UserReq{
			AdminUserId: uint64(this.AdminId()),
		},
	})
	if list != nil && len(list.Data.List) > 0 {
		this.Data["dbList"] = list.Data.List
	} else {
		this.Data["dbList"] = []maps.Map{}
	}

	fmt.Println("params==", params)
	fmt.Println("list==", list)
	if params.Json {
		this.Success()
	}
	this.Show()
}
