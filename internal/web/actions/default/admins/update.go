package admins

import (
	"encoding/json"
	"fmt"

	"github.com/1uLang/zhiannet-api/common/server/edge_admins_server"
	hids_user_model "github.com/1uLang/zhiannet-api/hids/model/user"
	hids_user_server "github.com/1uLang/zhiannet-api/hids/server/user"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	nc_req "github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/systemconfigs"
	"github.com/dlclark/regexp2"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/xlzd/gotp"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "update")
}

func (this *UpdateAction) RunGet(params struct {
	AdminId int64
}) {
	adminResp, err := this.RPC().AdminRPC().FindEnabledAdmin(this.AdminContext(), &pb.FindEnabledAdminRequest{AdminId: params.AdminId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	admin := adminResp.Admin
	if admin == nil {
		this.NotFound("admin", params.AdminId)
		return
	}

	// OTP认证
	otpLoginIsOn := false
	if admin.OtpLogin != nil {
		otpLoginIsOn = admin.OtpLogin.IsOn
	}

	// AccessKey数量
	countAccessKeyResp, err := this.RPC().UserAccessKeyRPC().CountAllEnabledUserAccessKeys(this.AdminContext(), &pb.CountAllEnabledUserAccessKeysRequest{AdminId: params.AdminId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	countAccessKeys := countAccessKeyResp.Count

	this.Data["admin"] = maps.Map{
		"id":              admin.Id,
		"fullname":        admin.Fullname,
		"username":        admin.Username,
		"isOn":            admin.IsOn,
		"isSuper":         admin.IsSuper,
		"canLogin":        admin.CanLogin,
		"otpLoginIsOn":    otpLoginIsOn,
		"countAccessKeys": countAccessKeys,
	}

	// 权限
	moduleMaps := configloaders.AllModuleMaps()
	for _, m := range moduleMaps {
		code := m.GetString("code")
		isChecked := false
		for _, module := range admin.Modules {
			if module.Code == code {
				isChecked = true
				break
			}
		}
		m["isChecked"] = isChecked
	}
	this.Data["modules"] = moduleMaps

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	AdminId int64

	Fullname    string
	Username    string
	Pass1       string
	Pass2       string
	ModuleCodes []string
	IsOn        bool
	IsSuper     bool
	CanLogin    bool

	// OTP
	OtpOn bool

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("修改系统用户 %d", params.AdminId)

	params.Must.
		Field("fullname", params.Fullname).
		Require("请输入系统用户全名")

	params.Must.
		Field("username", params.Username).
		Require("请输入登录用户名").
		Match(`^[a-zA-Z0-9_]+$`, "用户名中只能包含英文、数字或下划线")

	existsResp, err := this.RPC().AdminRPC().CheckAdminUsername(this.AdminContext(), &pb.CheckAdminUsernameRequest{
		AdminId:  params.AdminId,
		Username: params.Username,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if existsResp.Exists {
		this.FailField("username", "此用户名已经被别的系统用户使用，请换一个")
	}

	var editPwd bool
	if len(params.Pass1) > 0 {
		params.Must.
			Field("pass1", params.Pass1).
			Require("请输入登录密码").
			Field("pass2", params.Pass2).
			Require("请输入确认登录密码")
		if params.Pass1 != params.Pass2 {
			this.FailField("pass2", "两次输入的密码不一致")
		}
		reg, err := regexp2.Compile(
			`^(?![A-z0-9]+$)(?=.[^%&',;=?$\x22])(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).{8,30}$`, 0)
		if err != nil {
			this.FailField("pass1", "密码格式不正确")
		}
		if match, err := reg.FindStringMatch(params.Pass1); err != nil || match == nil {
			this.FailField("pass1", "密码格式不正确")
		}
		editPwd = true
	}

	// 修改nc密码
	if params.AdminId != 1 {
		pt, err := model.GetUsername(params.AdminId, 1)
		if err != nil {
			this.ErrorPage(err)
		}
		ncName, pp, err := nc_req.ParseToken(pt)
		if err != nil {
			this.ErrorPage(err)
		}
		if pp != params.Pass2 && params.Pass2 != ""{
			err = nc_req.UpdateUserPassword(params.Pass2, ncName)
			if err != nil {
				this.ErrorPage(err)
			}
			token := &model.LoginReq{
				User:     ncName,
				Password: params.Pass2,
			}
			ncToken := nc_req.GenerateToken(token)
			err = model.StoreNCToken(ncName, ncToken)
			if err != nil {
				this.ErrorPage(err)
			}
		}
	}

	modules := []*systemconfigs.AdminModule{}
	for _, code := range params.ModuleCodes {
		modules = append(modules, &systemconfigs.AdminModule{
			Code:     code,
			AllowAll: true,
			Actions:  nil, // TODO 后期再开放细粒度控制
		})
	}
	modulesJSON, err := json.Marshal(modules)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = this.RPC().AdminRPC().UpdateAdmin(this.AdminContext(), &pb.UpdateAdminRequest{
		AdminId:     params.AdminId,
		Username:    params.Username,
		Password:    params.Pass1,
		Fullname:    params.Fullname,
		ModulesJSON: modulesJSON,
		IsSuper:     params.IsSuper,
		IsOn:        params.IsOn,
		CanLogin:    params.CanLogin,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 修改OTP
	otpLoginResp, err := this.RPC().LoginRPC().FindEnabledLogin(this.AdminContext(), &pb.FindEnabledLoginRequest{
		AdminId: params.AdminId,
		Type:    "otp",
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	{
		otpLogin := otpLoginResp.Login
		if params.OtpOn {
			if otpLogin == nil {
				otpLogin = &pb.Login{
					Id:   0,
					Type: "otp",
					ParamsJSON: maps.Map{
						"secret": gotp.RandomSecret(16), // TODO 改成可以设置secret长度
					}.AsJSON(),
					IsOn:    true,
					AdminId: params.AdminId,
					UserId:  0,
				}
			} else {
				// 如果已经有了，就覆盖，这样可以保留既有的参数
				otpLogin.IsOn = true
			}

			_, err = this.RPC().LoginRPC().UpdateLogin(this.AdminContext(), &pb.UpdateLoginRequest{Login: otpLogin})
			if err != nil {
				this.ErrorPage(err)
				return
			}
		} else {
			_, err = this.RPC().LoginRPC().UpdateLogin(this.AdminContext(), &pb.UpdateLoginRequest{Login: &pb.Login{
				Type:    "otp",
				IsOn:    false,
				AdminId: params.AdminId,
			}})
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}
		//判断是否拥有了主机防护功能  有 则对应创建该用户
		var hasHids bool
		for _, code := range params.ModuleCodes {
			if code == configloaders.AdminModuleCodeHids {
				hasHids = true
				break
			}
		}
		if !configloaders.AllowModule(this.AdminId(), configloaders.AdminModuleCodeHids) && hasHids {

			err = hids.InitAPIServer()
			if err != nil {
				this.ErrorPage(fmt.Errorf("主机防护组件初始化失败：%v", err))
				return
			}
			_, err = hids_user_server.Add(&hids_user_model.AddReq{UserName: params.Username, Password: "dengbao-" + params.Username, Role: 3})
			if err != nil {
				this.ErrorPage(fmt.Errorf("主机防护组件同步信息失败：%v", err))
				return
			}
		}

		// 通知更改
		err = configloaders.NotifyAdminModuleMappingChange()
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if editPwd {
			//更新密码修改时间
			edge_admins_server.UpdatePwdAt(uint64(params.AdminId))
		}
		this.Success()
	}
}
