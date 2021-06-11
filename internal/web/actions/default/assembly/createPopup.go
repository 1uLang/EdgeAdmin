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
	Type   int

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入节点名称")

	params.Must.
		Field("addr", params.Addr).
		Require("请输入地址")

	if params.Port <= 0 {
		this.Fail("请选择端口")
	}

	switch params.Type {
	case 0: //ddos防火墙
		params.Must.
			Field("key", params.Key).
			Require("请输入username")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入password")
	case 1: //云防火墙
		params.Must.
			Field("key", params.Key).
			Require("请输入key")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入secret")
	case 2: //主机防护
		params.Must.
			Field("key", params.Key).
			Require("请输入appid")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入secret")
	case 3: //web漏扫
		params.Must.
			Field("key", params.Key).
			Require("请输入XAuth")
	case 4: //主机漏扫
		params.Must.
			Field("key", params.Key).
			Require("请输入accessKey")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入secretKey")

	}
	req := &subassemblynode_model.Subassemblynode{
		Name:   params.Name,
		Addr:   params.Addr,
		Port:   params.Port,
		Idc:    params.Idc,
		Type:   params.Type,
		Key:    params.Key,
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
