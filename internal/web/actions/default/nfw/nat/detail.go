package nat

import (
	"github.com/1uLang/zhiannet-api/opnsense/server/nat"
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
	info, err := nat.GetNat1To1Info(&nat.InfoReq{
		Id:     params.Id,
		NodeId: params.NodeId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//接口
	this.Data["interface"] = info.Interface
	this.Data["type"] = info.Type
	this.Data["external"] = info.External
	//this.Data["src"] = info.Src
	this.Data["src"] = ""
	this.Data["dst"] = info.Dst
	this.Data["dstmask"] = info.Dstmask
	this.Data["srcmask"] = info.Srcmask
	this.Data["descr"] = info.Descr

	//
	//// 创建日志
	//defer threateLog(oplogs.LevelInfo, "创建高防IP %d", params.Addr)

	this.Success()
}
