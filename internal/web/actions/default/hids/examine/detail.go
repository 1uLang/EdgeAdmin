package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	examine_server "github.com/1uLang/zhiannet-api/hids/server/examine"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) RunGet(params struct {
	MacCode string

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")

	if err := hids.InitAPIServer(); err != nil {
		this.ErrorPage(err)
		return
	}
	info, err := examine_server.Details(&examine.DetailsReq{
		MacCode:     params.MacCode,
		AdminUserId: uint64(this.AdminId()),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	list, err := examine_server.List(&examine.SearchReq{AdminUserId: uint64(this.AdminId()), Type: -1, Score: -1, State: -1, MacCode: params.MacCode})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["details"] = info
	this.Data["otherDetails"] = list.ServerExamineResultInfoList[0]

	this.Show()
}
