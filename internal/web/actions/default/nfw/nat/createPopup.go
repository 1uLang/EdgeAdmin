package nat

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/server/nat"
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
	NodeId    uint64
	Interface string
	Type      string
	External  string
	Src       string
	Srcmask   string
	Dsts      string
	Dst       string
	Dstmask   string
	Descr     string
	Id        string

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	this.Data["params"] = params
	//this.Show()
	this.Success()
	return
	params.Must.
		Field("nodeId", params.NodeId).
		Require("没有选择节点").
		Field("external", params.External).
		Require("请输入外部网络").
		Field("src", params.Src).
		Require("请输入源")

	if params.Dsts == "" { //目标
		params.Dsts = params.Dst
	}
	//this.Data["params"] = params
	//this.Success()
	//return
	params.Must.
		Field("dsts", params.Dsts).
		Require("请选择或输入源")

	data := &nat.SaveNat1To1Req{
		NodeId:    params.NodeId,
		Interface: params.Interface,
		Type:      params.Type,
		External:  params.External,
		Src:       params.Src,
		Srcmask:   params.Srcmask,
		Dst:       params.Dsts,
		Dstmask:   params.Dstmask,
		Descr:     params.Descr,
		ID:        params.Id,
	}
	tips, err := nat.SaveNat1To1(data)
	fmt.Println("err==", err, "tips==", tips)
	if err != nil || len(tips) > 0 {
		this.ErrorPage(err)
		return
	}
	defer this.CreateLogInfo("编辑nat 1：1  %d", params.Id)
	//this.Data["tops"] = tops
	//this.Show()
	this.Success()
}
