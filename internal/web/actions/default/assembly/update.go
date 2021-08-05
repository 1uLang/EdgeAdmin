package assembly

import (
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
		"id":        info.Id,
		"name":      info.Name,
		"addr":      info.Addr,
		"port":      info.Port,
		"argeement": info.IsSsl,
		"type":      info.Type,
		"idc":       info.Idc,
		//"Status": info.Status,
		"State":  info.State,
		"key":    "",
		"secret": "",
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
	//Status       int
	State     int
	Key       string
	Secret    string
	Argeement int

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("修改节点 %d", params.Id)

	params.Must.
		Field("name", params.Name).
		Require("请输入节点名称")

	params.Must.
		Field("addr", params.Addr).
		Require("请输入地址").
		Match("(^(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\d|1\\d{2}"+
			"|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\"+
			"d|1\\d{2}|2[0-4]\\d|25[0-5]):([0-9]|[1-9]\\d|[1-9]\\d{2}|[1-9]\\d{3}|[1-5]\\d{4}|6[0-4]\\d{3}|65[0-4]\\d{2}|655[0-"+
			"2]\\d|6553[0-5])$)|([a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?)", "请输入正确的地址")

	//检测节点类型
	_, isExist := typeMap[params.AssemblyType]
	if !isExist {
		this.Fail("请选择节点类型")
	}
	//检测数据中心
	_, isExist = idcMap[params.IdcId]
	if !isExist {
		this.Fail("请选择数据中心")
	}

	_, err := subassemblynode_server.Edit(&subassemblynode.Subassemblynode{
		Id:     params.Id,
		Name:   params.Name,
		Addr:   params.Addr,
		Port:   params.Port,
		Type:   params.AssemblyType,
		Idc:    params.IdcId,
		State:  1,
		Key:    params.Key,
		Secret: params.Secret,
		IsSsl:  params.Argeement,
	})
	if err != nil {
		this.Fail(err.Error())
		return
	}

	this.Success()
}
