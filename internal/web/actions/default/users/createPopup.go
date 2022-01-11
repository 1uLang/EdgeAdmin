package users

import (
	"github.com/1uLang/zhiannet-api/common/model/edge_logins"
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_users_server"
	"github.com/1uLang/zhiannet-api/nextcloud"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	nc_req "github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/utils/numberutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
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
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Username  string
	Pass1     string
	Pass2     string
	Fullname  string
	Mobile    string
	Tel       string
	Email     string
	Remark    string
	ClusterId int64
	OtpOn     bool
	ChannelId uint64

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("username", params.Username).
		Require("请输入用户名").
		Match(`^[a-zA-Z0-9_]+$`, "用户名中只能含有英文、数字和下划线")

	checkUsernameResp, err := this.RPC().UserRPC().CheckUserUsername(this.AdminContext(), &pb.CheckUserUsernameRequest{
		UserId:   0,
		Username: params.Username,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if checkUsernameResp.Exists {
		this.FailField("username", "此用户名已经被占用，请换一个")
	}

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

	params.Must.
		Field("fullname", params.Fullname).
		Require("请输入全名")

	if params.ClusterId <= 0 {
		this.Fail("请选择关联集群")
	}

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
	if nextcloud.UseDatabackup {
		// 创建nextcloud账号，并写入数据库
		adminToken := nc_req.GetAdminToken()
		// userPwd := `adminAd#@2021`
		userPwd := params.Pass2
		err = nc_req.CreateUserV2(adminToken, params.Username, userPwd)
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
		err = model.StoreNCToken(params.Username, ncToken)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	createResp, err := this.RPC().UserRPC().CreateUser(this.AdminContext(), &pb.CreateUserRequest{
		Username:      params.Username,
		Password:      params.Pass1,
		Fullname:      params.Fullname,
		Mobile:        params.Mobile,
		Tel:           params.Tel,
		Email:         params.Email,
		Remark:        params.Remark,
		Source:        "admin:" + numberutils.FormatInt64(this.AdminId()),
		NodeClusterId: params.ClusterId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.ChannelId > 0 { //修改渠道ID
		edge_users_server.UpdateChannel(uint64(createResp.UserId), params.ChannelId)
	}
	defer this.CreateLogInfo("创建用户 %d", createResp.UserId)

	if nextcloud.UseDatabackup {
		// 用户账号和nextcloud账号进行关联
		// 因为用户名是唯一的，所以加入用户名字段，减少脏数据的产生
		err = model.BindNCTokenAndUID(params.Username, createResp.UserId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	//更新密码修改时间
	edge_users_server.UpdatePwdAt(uint64(createResp.UserId))
	//otp
	if params.OtpOn {
		otpLogin := &edge_logins.EdgeLogins{
			Id:   0,
			Type: "otp",
			Params: string(maps.Map{
				"secret": gotp.RandomSecret(16), // TODO 改成可以设置secret长度
			}.AsJSON()),
			IsOn:    1,
			AdminId: 0,
			UserId:  uint64(createResp.UserId),
			State:   1,
		}

		_, err = edge_logins_server.Save(otpLogin)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	this.Success()
}
