package assembly

import (
	subassemblynode_model "github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/common/server/subassemblynode"
	nc_req "github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"

	audit_request "github.com/1uLang/zhiannet-api/audit/request"
	ddos_request "github.com/1uLang/zhiannet-api/ddos/request"
	hids_request "github.com/1uLang/zhiannet-api/hids/server"
	maltrail_request "github.com/1uLang/zhiannet-api/maltrail/request"
	nessus_request "github.com/1uLang/zhiannet-api/nessus/server"
	term_request "github.com/1uLang/zhiannet-api/next-terminal/server"
	awvs_request "github.com/1uLang/zhiannet-api/nextcloud/request"
	nextcloud_request "github.com/1uLang/zhiannet-api/nextcloud/request"
	opnsense_request "github.com/1uLang/zhiannet-api/opnsense/request"
	teaweb_request "github.com/1uLang/zhiannet-api/resmon/request"
	wazuh_request "github.com/1uLang/zhiannet-api/wazuh/server"
	zstack_request "github.com/1uLang/zhiannet-api/zstack/request"
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
	Idc          string
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
		Require("请输入节点名称").
		Field("idc", params.Idc).
		Require("请输入数据中心称")

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
	case 6, 7, 10, 11, 12: //审计系统 堡垒机 云底座，apt,wazuh
		params.Must.
			Field("key", params.Key).
			Require("请输入username")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入password")
	case 8: //数据备份
		params.Must.
			Field("key", params.Key).
			Require("请输入username")
		params.Must.
			Field("secret", params.Secret).
			Require("请输入password")

	case 9: //节点监控
		params.Must.
			Field("key", params.Key).
			Require("请输入密钥")
	}
	req := &subassemblynode_model.Subassemblynode{
		Name:   params.Name,
		Addr:   params.Addr,
		Port:   params.Port,
		Idc:    params.Idc,
		Type:   params.AssemblyType,
		State:  1,
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

	// 创建成功则和nextcloud关联
	if params.AssemblyType == 8 {
		err := nc_req.ConnNextcloudWithAdmin(params.Key, params.Secret)
		if err != nil {
			this.Fail(err.Error())
			return
		}
	}
	go this.Check(params.AssemblyType)
	this.Success()
}

func (this *CreatePopupAction) Check(AssemblyType int) {
	switch AssemblyType {
	case 1: //DDoS防护
		check := new(ddos_request.LoginReq)
		check.Run()
	case 2: //云防火墙
		check := new(opnsense_request.ApiKey)
		check.Run()
	case 3: //主机漏洞扫描
		check := new(nessus_request.CheckRequest)
		check.Run()

	case 4: //WEB漏洞扫描
		check := new(awvs_request.CheckRequest)
		check.Run()

	case 5: //主机防护
		check := new(hids_request.CheckRequest)
		check.Run()
	case 6: //安全审计
		check := new(audit_request.LoginReq)
		check.Run()
	case 7: //堡垒机
		check := new(term_request.CheckRequest)
		check.Run()
	case 8: //数据备份
		check := new(nextcloud_request.CheckRequest)
		check.Run()

	case 9: //节点监控
		check := new(teaweb_request.CheckRequest)
		check.Run()
	case 10: //云底座
		check := new(zstack_request.LoginReq)
		check.Run()
	case 11: //apt检测
		check := new(maltrail_request.LoginReq)
		check.Run()
	case 12: //wazuh检测
		check := new(wazuh_request.CheckRequest)
		check.Run()
	}
}
