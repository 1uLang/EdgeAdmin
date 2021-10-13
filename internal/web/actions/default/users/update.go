package users

import (
	"github.com/1uLang/zhiannet-api/common/model/edge_logins"
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_users_server"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	nc_req "github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/users/userutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
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
	UserId int64
}) {
	err := userutils.InitUser(this.Parent(), params.UserId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	//userResp, err := this.RPC().UserRPC().FindEnabledUser(this.AdminContext(), &pb.FindEnabledUserRequest{UserId: params.UserId})

	user, err := edge_users_server.GetUserInfo(uint64(params.UserId))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//user := userResp.User
	//if user == nil {
	//	this.NotFound("user", params.UserId)
	//	return
	//}

	// AccessKey数量
	countAccessKeyResp, err := this.RPC().UserAccessKeyRPC().CountAllEnabledUserAccessKeys(this.AdminContext(), &pb.CountAllEnabledUserAccessKeysRequest{UserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	countAccessKeys := countAccessKeyResp.Count
	//otp
	var otpLogin bool
	info, err := edge_logins_server.GetInfoByUid(uint64(params.UserId))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if info != nil && info.IsOn == 1 {
		otpLogin = true
	}

	this.Data["user"] = maps.Map{
		"id":              user.ID,
		"username":        user.Username,
		"fullname":        user.Fullname,
		"email":           user.Email,
		"tel":             user.Tel,
		"remark":          user.Remark,
		"mobile":          user.Mobile,
		"isOn":            user.Ison,
		"countAccessKeys": countAccessKeys,
		"otpLoginIsOn":    otpLogin,
		"channelId":       user.ChannelId,
	}

	this.Data["clusterId"] = 0
	if user.Clusterid > 0 {
		this.Data["clusterId"] = user.Clusterid
	}

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	UserId    int64
	Username  string
	Pass1     string
	Pass2     string
	Fullname  string
	Mobile    string
	Tel       string
	Email     string
	Remark    string
	IsOn      bool
	ClusterId int64
	OtpOn     bool
	ChannelId uint64

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("修改用户 %d", params.UserId)

	params.Must.
		Field("username", params.Username).
		Require("请输入用户名").
		Match(`^[a-zA-Z0-9_]+$`, "用户名中只能含有英文、数字和下划线")

	checkUsernameResp, err := this.RPC().UserRPC().CheckUserUsername(this.AdminContext(), &pb.CheckUserUsernameRequest{
		UserId:   params.UserId,
		Username: params.Username,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if checkUsernameResp.Exists {
		this.FailField("username", "此用户名已经被占用，请换一个")
	}

	var editPwd bool
	if len(params.Pass1) > 0 {
		params.Must.
			Field("pass1", params.Pass1).
			Require("请输入密码").
			Field("pass2", params.Pass2).
			Require("请再次输入确认密码").
			Equal(params.Pass1, "两次输入的密码不一致")
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

	params.Must.
		Field("fullname", params.Fullname).
		Require("请输入全名")

	if len(params.Mobile) > 0 {
		params.Must.
			Field("mobile", params.Mobile).
			Mobile("请输入正确的手机号")
	}
	if len(params.Email) > 0 {
		params.Must.
			Field("email", params.Email).
			Email("请输入正确的电子邮箱")
	}

	_, err = this.RPC().UserRPC().UpdateUser(this.AdminContext(), &pb.UpdateUserRequest{
		UserId:        params.UserId,
		Username:      params.Username,
		Password:      params.Pass1,
		Fullname:      params.Fullname,
		Mobile:        params.Mobile,
		Tel:           params.Tel,
		Email:         params.Email,
		Remark:        params.Remark,
		IsOn:          params.IsOn,
		NodeClusterId: params.ClusterId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if editPwd {
		//更新密码修改时间
		edge_users_server.UpdatePwdAt(uint64(params.UserId))
	}

	//if params.ChannelId > 0 { //修改渠道ID
	edge_users_server.UpdateChannel(uint64(params.UserId), params.ChannelId)
	//}

	otpLogin, err := edge_logins_server.GetInfoByUid(uint64(params.UserId))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.OtpOn {
		if otpLogin == nil || otpLogin.Id == 0 {
			otpLogin = &edge_logins.EdgeLogins{
				Id:   0,
				Type: "otp",
				Params: string(maps.Map{
					"secret": gotp.RandomSecret(16), // TODO 改成可以设置secret长度
				}.AsJSON()),
				IsOn:    1,
				AdminId: 0,
				UserId:  uint64(params.UserId),
				State:   1,
			}
		} else {
			// 如果已经有了，就覆盖，这样可以保留既有的参数
			otpLogin.IsOn = 1
		}

		_, err = edge_logins_server.Save(otpLogin)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	} else {
		//fmt.Println("otp=====", otpLogin)
		if otpLogin != nil && otpLogin.Id > 0 {
			_, err = edge_logins_server.UpdateOpt(uint64(otpLogin.Id), 0)
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}

	}

	// 修改nc密码
	pt, err := model.GetUsername(params.UserId, 0)
	if err != nil {
		this.ErrorPage(err)
	}
	ncName, pp, err := nc_req.ParseToken(pt)
	if err != nil {
		this.ErrorPage(err)
	}
	if pp != params.Pass2 && params.Pass2 != "" {
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

	this.Success()
}
