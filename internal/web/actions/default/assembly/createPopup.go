package assembly

import (
	subassemblynode_model "github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/common/server/subassemblynode"
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
	Name   string
	Addr   string
	Port   int64
	Idc    int
	Key    string
	Secret string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	req := &subassemblynode_model.Subassemblynode{
		Name:   params.Name,
		Addr:   params.Addr,
		Port:   params.Port,
		Idc:    params.Idc,
		Secret: params.Secret,
	}
	id, err := subassemblynode.Add(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	defer this.CreateLogInfo("创建节点 %d", id)

	this.Success()
}
