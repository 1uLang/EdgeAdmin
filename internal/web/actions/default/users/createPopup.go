package users

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/model/audit_user_relation"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/user"
	"github.com/TeaOSLab/EdgeAdmin/internal/utils/numberutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
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

	params.Must.
		Field("pass1", params.Pass1).
		Match(`^[a-zA-Z0-9_]+$`, "用户名中只能含有英文、数字和下划线")

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

	// 创建nextcloud账号，并写入数据库
	adminToken := request.GetAdminToken()
	userPwd := `adminAd#@2021`
	err = request.CreateUser(adminToken, params.Username, userPwd)
	if err != nil {
		log.Println("createUser", err.Error())
		this.ErrorPage(err)
		return
	}
	// 生成token
	gtReq := &model.LoginReq{
		User:     params.Username,
		Password: userPwd,
	}
	ncToken := request.GenerateToken(gtReq)
	// 写入数据库
	err = model.StoreNCToken(params.Username, ncToken)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	//创建审计系统的账号
	//{
	auditResp, auditErr := user.AddUser(&user.AddUserReq{
		User:        &request.UserReq{AdminUserId: uint64(this.AdminId())},
		Email:       params.Email,
		IsAdmin:     1,
		NickName:    params.Username,
		Opt:         1,
		Password:    params.Pass1,
		Phonenumber: params.Mobile,
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

	//}

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
	defer this.CreateLogInfo("创建用户 %d", createResp.UserId)

	//关联账号
	_, err = audit_user_relation.Add(&audit_user_relation.AuditReq{
		UserId:      uint64(createResp.UserId),
		AuditUserId: uint64(auditResp.Data.Id),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 用户账号和nextcloud账号进行关联
	// 因为用户名是唯一的，所以加入用户名字段，减少脏数据的产生
	err = model.BindNCTokenAndUID(params.Username, createResp.UserId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
