package assets

import (
	"fmt"
	gateway_model "github.com/1uLang/zhiannet-api/next-terminal/model/access_gateway"
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	cert_model "github.com/1uLang/zhiannet-api/next-terminal/model/cert"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "fortcloud", "index")
}

func (this *IndexAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}
func (this *IndexAction) RunGet(params struct {
	PageSize  int
	PageNo    int
	PageState int
	Asset     string

	Must *actions.Must
}) {
	defer this.Show()

	this.Data["assets"] = []int64{}
	this.Data["certs"] = []int64{}
	this.Data["gateways"] = []int64{}

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	list, _, err := req.Assets.List(&asset_model.ListReq{AdminUserId: this.AdminId()})
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	certs, _, err := req.Cert.List(&cert_model.ListReq{AdminUserId: uint64(this.AdminId())})
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	gateways, err := req.GateWay.GetAll(&gateway_model.GetAllReq{AdminUserId: this.AdminId()})
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	for k, v := range list {
		item := v.(map[string]interface{})
		item["auth"] = strings.HasPrefix(item["tags"].(string), fmt.Sprintf("admin_%v", this.AdminId()))
		list[k] = item
	}
	this.Data["assets"] = list
	this.Data["certs"] = certs
	this.Data["gateways"] = gateways
}

func (this *IndexAction) RunPost(params struct {
	HostName    string
	Ip          string
	Type        string
	Description string
	Password    string
	Port        int
	Protocol    string
	Username    string
	CertId      string
	Gateway     string
	Must        *actions.Must
}) {

	params.Must.
		Field("hostName", params.HostName).
		Require("请输入主机名").
		Field("protocol", params.Protocol).
		Require("请选择接入协议").
		Field("port", params.Port).
		Require("请输入端口号").
		Field("type", params.Type).
		Require("请选择账户类型").
		Field("ip", params.Ip).
		Require("请输入主机地址").
		Match("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", "请输入正确的主机地址")

	if params.Type == "custom" {
		params.Must.
			Field("username", params.Username).
			Require("请输入授权账号").
			Field("protocol", params.Protocol).
			Require("请输入密码")
	} else if params.Type == "on_custom" { //免密登录
		params.Type = "custom"
	} else {
		params.Must.
			Field("certId", params.CertId).
			Require("请选择授权凭证")
	}
	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	args := &asset_model.CreateReq{}
	args.Name = params.HostName
	args.IP = params.Ip
	args.AccountType = params.Type
	args.Description = params.Description
	args.Password = params.Password
	args.Port = params.Port
	args.Protocol = params.Protocol
	args.Username = params.Username
	args.CredentialId = params.CertId
	args.AccessGatewayId = params.Gateway
	args.AdminUserId = uint64(this.AdminId())
	err = req.Assets.Create(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 新增资产:[%v]成功", params.HostName)
	this.Success()
}
