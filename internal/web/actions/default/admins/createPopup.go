package admins

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/model/audit_user_relation"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/user"
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

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{}) {
	this.Data["modules"] = configloaders.AllModuleMaps()
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Fullname    string
	Username    string
	Pass1       string
	Pass2       string
	ModuleCodes []string
	IsSuper     bool
	CanLogin    bool

	// OTP
	OtpOn bool

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("fullname", params.Fullname).
		Require("请输入系统用户全名")

	params.Must.
		Field("username", params.Username).
		Require("请输入登录用户名").
		Match(`^[a-zA-Z0-9_]+$`, "用户名中只能包含英文、数字或下划线")

	existsResp, err := this.RPC().AdminRPC().CheckAdminUsername(this.AdminContext(), &pb.CheckAdminUsernameRequest{
		AdminId:  0,
		Username: params.Username,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if existsResp.Exists {
		this.FailField("username", "此用户名已经被别的系统用户使用，请换一个")
	}

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

	// 创建nextcloud账号，并写入数据库
	adminToken := nc_req.GetAdminToken()
	userPwd := `adminAd#@2021`
	un := "admin_" + params.Username
	err = nc_req.CreateUser(adminToken, un, userPwd)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 生成token
	gtReq := &model.LoginReq{
		User:     params.Username,
		Password: userPwd,
	}
	ncToken := nc_req.GenerateToken(gtReq)
	// 写入数据库
	err = model.StoreNCToken(params.Username, ncToken, 1)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	//创建审计系统账号
	auditResp, auditErr := user.AddUser(&user.AddUserReq{
		User:        &request.UserReq{AdminUserId: uint64(this.AdminId())},
		Email:       "",
		IsAdmin:     1,
		NickName:    params.Username,
		Opt:         1,
		Password:    params.Pass1,
		Phonenumber: "",
		RoleIds:     []uint64{},
		RoleName:    "平台管理员",
		Sex:         1,
		Status:      1,
		UserName:    params.Username,
	})
	if auditErr != nil || auditResp == nil {
		this.ErrorPage(fmt.Errorf("创建账号失败"))
		return
	}
	if auditResp.Code != 0 {
		this.ErrorPage(fmt.Errorf(auditResp.Msg))
		return
	}

	createResp, err := this.RPC().AdminRPC().CreateAdmin(this.AdminContext(), &pb.CreateAdminRequest{
		Username:    params.Username,
		Password:    params.Pass1,
		Fullname:    params.Fullname,
		ModulesJSON: modulesJSON,
		IsSuper:     params.IsSuper,
		CanLogin:    params.CanLogin,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// OTP
	if params.OtpOn {
		_, err = this.RPC().LoginRPC().UpdateLogin(this.AdminContext(), &pb.UpdateLoginRequest{Login: &pb.Login{
			Id:   0,
			Type: "otp",
			ParamsJSON: maps.Map{
				"secret": gotp.RandomSecret(16), // TODO 改成可以设置secret长度
			}.AsJSON(),
			IsOn:    true,
			AdminId: createResp.AdminId,
			UserId:  0,
		}})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	//关联审计系统账号
	_, err = audit_user_relation.Add(&audit_user_relation.AuditReq{
		AdminUserId: uint64(createResp.AdminId),
		AuditUserId: uint64(auditResp.Data.Id),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	defer this.CreateLogInfo("创建系统用户 %d", createResp.AdminId)

	// 用户账号和nextcloud账号进行关联
	// 因为用户名是唯一的，所以加入用户名字段，减少脏数据的产生
	err = model.BindNCTokenAndUID(un, createResp.AdminId, 1)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	//判断是否拥有了主机防护功能  有 则对应创建该用户
	var hasHids bool
	for _, code := range params.ModuleCodes {
		if code == configloaders.AdminModuleCodeHids {
			hasHids = true
			break
		}
	}
	if hasHids {
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

	this.Success()
}
