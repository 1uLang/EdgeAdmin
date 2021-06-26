package acl

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/server/acl"
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
	NodeId     uint64
	Interface  string
	Type       string
	Direction  string
	Ipprotocol string
	Protocol   string
	Src        string
	Srcinput   string
	Srcmask    string
	Dst        string
	Dstinput   string
	Dstmask    string
	Descr      string
	Id         string

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
		//Field("external", params.External).
		//Require("请输入外部网络").
		Field("src", params.Src).
		Require("请输入源")

	if params.Dst == "" { //目标
		params.Dst = params.Dstinput
	}
	//this.Data["params"] = params
	//this.Success()
	//return
	params.Must.
		Field("dsts", params.Dst).
		Require("请选择或输入源")

	data := &acl.SaveAclReq{
		NodeId:     params.NodeId,
		Interface:  params.Interface,
		Type:       params.Type,
		Direction:  params.Direction,
		Ipprotocol: params.Ipprotocol,
		Protocol:   params.Protocol,
		Src:        params.Src,
		Srcmask:    params.Srcmask,
		Dst:        params.Dst,
		Dstmask:    params.Dstmask,
		Descr:      params.Descr,
		ID:         params.Id,
	}
	tips, err := acl.SaveAcl(data)
	fmt.Println("err==", err, "tips==", tips)
	if err != nil || len(tips) > 0 {
		this.ErrorPage(err)
	}
	defer this.CreateLogInfo("编辑acl  %d", params.Id)
	//this.Data["tops"] = tops
	//this.Show()
	this.Success()
}
