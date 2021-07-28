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
	this.Data["state"] = 1
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name         string
	Addr         string
	Port         int64
	IdcId        int
	Key          string
	Secret       string
	AssemblyType int
	Argeement    int
	State        int

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
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
	switch params.AssemblyType {
	case 1: //ddos防火墙
		params.Must.
			Field("key", params.Key).
			Require("请输入username")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入password")
	case 2: //云防火墙
		params.Must.
			Field("key", params.Key).
			Require("请输入username")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入password")
	case 3: //主机漏洞扫描
		params.Must.
			Field("key", params.Key).
			Require("请输入accessKey")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入secretKey")
	case 4: //web漏扫
		params.Must.
			Field("key", params.Key).
			Require("请输入XAuth")
	case 5: //主机防护
		params.Must.
			Field("key", params.Key).
			Require("请输入appid")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入secret")
	case 7: //堡垒机
		params.Must.
			Field("key", params.Key).
			Require("请输入username")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入password")
	}
	req := &subassemblynode_model.Subassemblynode{
		Name:   params.Name,
		Addr:   params.Addr,
		Port:   params.Port,
		Idc:    params.IdcId,
		Type:   params.AssemblyType,
		State:  params.State,
		Key:    params.Key,
		Secret: params.Secret,
		IsSsl:  params.Argeement,
	}
	id, err := subassemblynode.Add(req)
	if err != nil {
		this.Fail(err.Error())
		return
	}
	defer this.CreateLogInfo("创建节点 %d", id)

	this.Success()
}
