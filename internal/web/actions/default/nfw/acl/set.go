package acl

import (
	"fmt"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server/acl"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/sirupsen/logrus"
)

type SetAction struct {
	actionutils.ParentAction
}

func (this *SetAction) Init() {
	this.Nav("", "", "")
}

func (this *SetAction) RunGet(params struct{}) {
	this.Show()
}

func (this *SetAction) RunPost(params struct {
	Id        string
	NodeId    uint64
	Status    string //状态 0禁用 1启用
	Interface string //类型  wan  lan

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("ID不能为空")
	params.Must.
		Field("node_id", params.NodeId).
		Require("节点不能为空")
	logrus.Info("nodeid", params.NodeId)
	res, err := opnsense_server.StartUpAcl(&opnsense_server.StartAclReq{
		ID:        params.Id,
		NodeId:    params.NodeId,
		Interface: params.Interface,
	})
	if err != nil || !res {
		this.ErrorPage(fmt.Errorf("操作失败"))
		return
	}

	//
	//// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "修改云防火墙acl规则状态 节点=%d id=%d", params.NodeId, params.Id)

	this.Success()
}
