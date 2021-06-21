package host

import (
	"fmt"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type SetAction struct {
	actionutils.ParentAction
}

func (this *SetAction) Init() {
	this.Nav("", "", "")
}

func (this *SetAction) RunGet(params struct {
	Addr   string
	NodeId uint64

	Must *actions.Must
}) {
	params.Must.
		Field("Addr", params.Addr).
		Require("请输入主机IP地址")

	if params.NodeId == 0 {
		this.ErrorPage(fmt.Errorf("请选择节点"))
		return
	}
	res, err := host_status_server.GetHostInfo(&host_status_server.HostGetReq{
		Addr:   params.Addr,
		NodeId: params.NodeId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	fmt.Println("-----------------", res.Ignore)
	this.Data["ignore"] = res.Ignore == "checked"
	this.Data["level"] = res.ParamSet
	this.Success()
}

func (this *SetAction) RunPost(params struct {
	Addr   string
	NodeId uint64
	Ignore bool //IP直通
	Set    int  //防护等级

	Must *actions.Must
}) {
	params.Must.
		Field("Addr", params.Addr).
		Require("请输入主机IP地址")

	if params.NodeId == 0 {
		this.ErrorPage(fmt.Errorf("请选择节点"))
		return
	}
	if params.Set < 0 || params.Set > 4 {
		this.ErrorPage(fmt.Errorf("防护策略参数无效"))
		return
	}
	_, err := host_status_server.SetHost(&host_status_server.HostSetReq{
		Addr:   params.Addr,
		NodeId: params.NodeId,
		Ignore: params.Ignore,
		//一下一个参数需统一
		ProtectSet: params.Set,
		FilterSet:  params.Set,
		SetTcp:     params.Set,
		SetUdp:     params.Set,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 创建日志
	//defer this.CreateLog(oplogs.LevelInfo, "创建高防IP %d", params.Addr)

	this.Success()
}
