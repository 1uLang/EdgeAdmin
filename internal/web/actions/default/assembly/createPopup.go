package assembly

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{}) {
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name       string
	Addr       string
	IdcId      int64
	AssemblyId int64

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入节点名称")
	params.Must.
		Field("addr", params.Addr).
		Require("请输入节点地址")
	params.Must.
		Field("idcId", params.IdcId).
		Require("请选择所属数据中心")
	params.Must.
		Field("assemblyType", params.AssemblyId).
		Require("请选择节点类型")

	defer this.CreateLogInfo("添加节点 %s 成功", params.Name)

	this.Success()
}
