package assembly

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	subassemblynode_server "github.com/1uLang/zhiannet-api/common/server/subassemblynode"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "update")
}

func (this *UpdateAction) RunGet(params struct {
	NodeId int64
}) {

	info, err := subassemblynode_server.GetNodeInfo(uint64(params.NodeId))
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if info == nil {
		this.NotFound("node id ", params.NodeId)
		return
	}

	this.Data["node"] = maps.Map{
		"id":     info.Id,
		"name":   info.Name,
		"addr":   info.Addr,
		"port":   info.Port,
		"type":   info.Type,
		"idc":    info.Idc,
		"Status": info.Status,
		"State":  info.State,
		"key":    info.Key,
		"secret": info.Secret,
	}

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	Id           uint64
	Name         string
	Addr         string
	Port         int64
	AssemblyType int
	IdcId        int
	Status       int
	State        int
	Key          string
	Secret       string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("修改节点 %d", params.Id)

	params.Must.
		Field("name", params.Name).
		Require("请输入节点名称")

	params.Must.
		Field("assemblyType", params.AssemblyType).
		Require("请选择节点类型")

	params.Must.
		Field("idcId", params.IdcId).
		Require("请选择数据中心")

	fmt.Println(params)

	_, err := subassemblynode_server.Edit(&subassemblynode.Subassemblynode{
		Id:     params.Id,
		Name:   params.Name,
		Addr:   params.Addr,
		Port:   params.Port,
		Type:   params.AssemblyType,
		Idc:    params.IdcId,
		State:  params.State,
		Key:    params.Key,
		Secret: params.Secret,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
