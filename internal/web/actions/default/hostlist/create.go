package hostlist

import (
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"github.com/1uLang/zhiannet-api/zstack/server/host_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{ NodeId uint64 }) {

	////ddos节点
	//ddos, err := host_status_server.GetDDoSNodeInfo(params.NodeId)
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	//this.Data["ddos"] = ddos.Name
	//this.Data["node"] = params.NodeId
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name        string
	SpecUuid    string
	ImageUuid   string
	DiskUuid    string
	NetworkUuid []string
	Must        *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入名称").
		Field("SpecUuid", params.SpecUuid).
		Require("请选择规格").
		Field("ImageUuid", params.ImageUuid).
		Require("请选择镜像").
		Field("DiskUuid", params.DiskUuid).
		Require("请选择云盘")

	if len(params.NetworkUuid) == 0 {
		this.Fail("请选择网络组")
	}

	res, err := host_server.CreateHost(&host.CreateHostReq{
		Params: host.ParamsHost{
			Name:                 params.Name,
			ImageUuid:            params.ImageUuid,
			L3NetworkUuids:       params.NetworkUuid,
			RootDiskOfferingUuid: params.DiskUuid,
			InstanceOfferingUuid: params.SpecUuid,
		},
	}, uint64(this.AdminId()))
	if err != nil {
		this.Fail(err.Error())
		return
	}
	if res.Error.Code != "" {
		this.Fail(res.Error.Description)
		return
	}
	this.Success()
}
