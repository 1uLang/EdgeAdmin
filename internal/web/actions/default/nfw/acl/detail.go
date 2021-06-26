package acl

import (
	"github.com/1uLang/zhiannet-api/opnsense/server/acl"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) Init() {
	this.Nav("", "", "")
}

func (this *DetailAction) RunGet(params struct {
	Id     string
	NodeId uint64

	Must *actions.Must
}) {
	params.Must.
		Field("nodeId", params.NodeId).
		Require("请选择节点")
	info, err := acl.GetAclInfo(&acl.InfoReq{
		ID:     params.Id,
		NodeId: params.NodeId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//接口
	this.Data["interface"] = info.Interface   //接口
	this.Data["disabled"] = info.Disabled     //状态 是否启用
	this.Data["type"] = info.Type             //操作
	this.Data["direction"] = info.Direction   //方向
	this.Data["ipprotocol"] = info.Ipprotocol //ip/tcp版本
	this.Data["protocol"] = info.Protocol     //协议
	this.Data["srcnot"] = info.Srcnot         //源反转
	this.Data["src"] = info.Src               //源
	this.Data["srcmask"] = info.Srcmask       //源掩码
	this.Data["dstnot"] = info.Dstnot         //目标反转
	this.Data["dst"] = info.Dst               //目标
	this.Data["dstmask"] = info.Dstmask       //目标掩码
	this.Data["descr"] = info.Descr           //描述

	//
	//// 创建日志
	//defer threateLog(oplogs.LevelInfo, "创建高防IP %d", params.Addr)

	this.Success()
}
