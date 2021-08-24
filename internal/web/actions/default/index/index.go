package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/server/edge_admins_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeAdmin/internal/const"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/rpc"
	"github.com/TeaOSLab/EdgeAdmin/internal/setup"
	"github.com/TeaOSLab/EdgeAdmin/internal/utils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
	stringutil "github.com/iwind/TeaGo/utils/string"
	"github.com/xlzd/gotp"
)

type IndexAction struct {
	actionutils.ParentAction
}

// 首页（登录页）

var TokenSalt = stringutil.Rand(32)

func (this *IndexAction) RunGet(params struct {
	From  string
	Token string

	Auth *helpers.UserShouldAuth
}) {
	// DEMO模式
	this.Data["isDemo"] = teaconst.IsDemoMode

	// 检查系统是否已经配置过
	if !setup.IsConfigured() {
		this.RedirectURL("/setup")
		return
	}

	// 已登录跳转到dashboard
	if params.Auth.IsUser() {
		this.RedirectURL("/dashboard")
		return
	}

	this.Data["isUser"] = false
	this.Data["menu"] = "signIn"

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	this.Data["token"] = stringutil.Md5(TokenSalt+timestamp) + timestamp
	this.Data["from"] = params.From

	uiConfig, err := configloaders.LoadAdminUIConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["systemName"] = uiConfig.AdminSystemName
	this.Data["showVersion"] = uiConfig.ShowVersion
	if len(uiConfig.Version) > 0 {
		this.Data["version"] = uiConfig.Version
	} else {
		this.Data["version"] = teaconst.Version
	}
	//this.Data["faviconFileId"] = config.FaviconFileId
	if params.Token != "" {
		this.Success()
	}
	this.Data["faviconFileId"] = uiConfig.FaviconFileId

	securityConfig, err := configloaders.LoadSecurityConfig()
	if err != nil {
		this.Data["rememberLogin"] = false
	} else {
		this.Data["rememberLogin"] = securityConfig.AllowRememberLogin
	}

	this.Show()
}

// RunPost 提交
func (this *IndexAction) RunPost(params struct {
	Token    string
	Username string
	Password string
	OtpCode  string
	Remember bool
	Must     *actions.Must
	Auth     *helpers.UserShouldAuth
	//CSRF     *actionutils.CSRF
}) {
	this.Data["from"] = ""

	//params.Must.
	//	Field("username", params.Username).
	//	Require("请输入用户名").
	//	Field("password", params.Password).
	//	Require("请输入密码")
	if params.Username == "" {
		this.FailField("username", "请输入用户名")
	}
	if params.Password == stringutil.Md5("") {
		this.FailField("password", "请输入密码")
	}

	// 检查token
	if len(params.Token) <= 32 {
		this.Fail("请通过登录页面登录")
	}
	timestampString := params.Token[32:]
	if stringutil.Md5(TokenSalt+timestampString) != params.Token[:32] {
		this.FailField("refresh", "登录页面已过期，请刷新后重试")
	}
	timestamp := types.Int64(timestampString)
	if timestamp < time.Now().Unix()-1800 {
		this.FailField("refresh", "登录页面已过期，请刷新后重试")
	}

	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		this.Fail("服务器出了点小问题：" + err.Error())
	}
	//登录限制检查
	if res, _ := edge_admins_server.LoginCheck(fmt.Sprintf("admin_%v", params.Username)); res {
		this.FailField("refresh", "账号已被锁定（请 30分钟后重试）")
	}

	resp, err := rpcClient.AdminRPC().LoginAdmin(rpcClient.Context(0), &pb.LoginAdminRequest{
		Username: params.Username,
		Password: params.Password,
	})

	if err != nil {
		err = dao.SharedLogDAO.CreateAdminLog(rpcClient.Context(0), oplogs.LevelError, this.Request.URL.Path, "登录时发生系统错误："+err.Error(), this.RequestRemoteIP())
		if err != nil {
			utils.PrintError(err)
		}

		actionutils.Fail(this, err)
	}

	if !resp.IsOk {
		err = dao.SharedLogDAO.CreateAdminLog(rpcClient.Context(0), oplogs.LevelWarn, this.Request.URL.Path, "登录失败，用户名："+params.Username, this.RequestRemoteIP())
		if err != nil {
			utils.PrintError(err)
		}
		info, err := edge_admins_server.GetUserInfoByName(params.Username)
		if err != nil {
			this.ErrorPage(err)
		}
		if (info != nil && info.Id > 0) && (info.State == 0 || info.Ison == 0) {
			this.Fail("当前账号被禁用")
		} else {
			//登录次数+1
			edge_admins_server.LoginErrIncr(fmt.Sprintf("admin_%v", params.Username))
			//num, _ := cache.GetInt(fmt.Sprintf("admin_%v", params.Username))
			this.Fail("登录失败，请重新登录")
			//this.Fail(fmt.Sprintf("请输入正确的用户名密码，您还可以尝试%v次，（账号将被临时锁定30分钟）", 5-num))
		}

	}

	// 检查OTP-*/
	otpLoginResp, err := this.RPC().LoginRPC().FindEnabledLogin(this.AdminContext(), &pb.FindEnabledLoginRequest{
		AdminId: resp.AdminId,
		Type:    "otp",
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if otpLoginResp.Login != nil && otpLoginResp.Login.IsOn {
		loginParams := maps.Map{}
		err = json.Unmarshal(otpLoginResp.Login.ParamsJSON, &loginParams)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		secret := loginParams.GetString("secret")
		if gotp.NewDefaultTOTP(secret).Now() != params.OtpCode {
			this.Fail("请输入正确的OTP动态密码")
		}
	}

	//密码过期检查
	//if res, _ := edge_admins_server.CheckPwdInvalid(uint64(resp.AdminId)); res {
	//	params.Auth.SetUpdatePwdToken(resp.AdminId)
	//	this.Data["from"] = "/updatePwd"
	//	this.Fail("密码已过期，请立即修改")
	//}
	//检测系统是否到期
	code, expire, err := checkExpire()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if expire {
		this.Data["from"] = "/renewal"
		this.Data["systemCode"] = code
		this.Fail("系统已到期，请立即续订")
	}

	adminId := resp.AdminId
	params.Auth.StoreAdmin(adminId, params.Remember)

	// 记录日志
	err = dao.SharedLogDAO.CreateAdminLog(rpcClient.Context(adminId), oplogs.LevelInfo, this.Request.URL.Path, "成功登录系统，用户名："+params.Username, this.RequestRemoteIP())
	if err != nil {
		utils.PrintError(err)
	}
	//记录登录成功30分钟
	cache.SetNx(fmt.Sprintf("login_success_adminid_%v", adminId), time.Minute*30)
	cache.DelKey(fmt.Sprintf("admin_%v", params.Username))
	this.Success()
}
