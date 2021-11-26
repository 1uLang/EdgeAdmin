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
	NodeId       uint64
	Interface    string
	Type         string
	Direction    string
	Ipprotocol   string
	Protocol     string
	Src          string
	Srcinput     string
	Srcmask      string
	Dst          string
	Dstinput     string
	Dstmask      string
	Descr        string
	Id           string
	Srcbeginport string
	Srcendport   string
	Dstbeginport string
	Dstendport   string
	Must         *actions.Must
	//CSRF *actionutils.CSRF
}) {
	//this.Data["params"] = params
	//this.Show()
	//this.Success()
	//return
	params.Must.
		Field("nodeId", params.NodeId).
		Require("没有选择节点").
		Field("src", params.Src).
		Field("srcinput", params.Srcinput).
		Require("请选择或输入源").
		Field("dstinput", params.Dstinput).
		Require("请选择或输入目标")

	data := &acl.SaveAclReq{
		NodeId:       params.NodeId,
		Interface:    params.Interface,
		Type:         params.Type,
		Direction:    params.Direction,
		Ipprotocol:   params.Ipprotocol,
		Protocol:     params.Protocol,
		Src:          params.Srcinput,
		Srcmask:      params.Srcmask,
		Dst:          params.Dstinput,
		Dstmask:      params.Dstmask,
		Descr:        params.Descr,
		ID:           params.Id,
		SrcBeginPort: params.Srcbeginport,
		SrcEndPort:   params.Srcendport,
		DstBeginPort: params.Dstbeginport,
		DstEndPort:   params.Dstendport,
	}
	tips, err := acl.SaveAcl(data)
	if err != nil {
		this.ErrorPage(err)
	}
	if len(tips) > 0 {
		err = fmt.Errorf(tips[0])
		this.ErrorPage(err)
	}
	defer this.CreateLogInfo("编辑acl  %d", params.Id)
	//this.Data["tops"] = tops
	//this.Show()
	this.Success()
}
