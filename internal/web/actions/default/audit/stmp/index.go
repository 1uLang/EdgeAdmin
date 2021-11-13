package stmp

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/email"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "audit", "index")
}

func (this *IndexAction) RunGet(params struct {
	PageNum  int
	PageSize int
}) {
	rsp, err := email.GetEmail(&email.EmailInfoReq{
		User: &request.UserReq{
			AdminUserId: uint64(this.AdminId()),
		},
	})
	this.Data["to"] = ""
	if err != nil || rsp == nil {
		//this.ErrorPage(fmt.Errorf("获取失败"))
		this.Data["host"] = ""
		this.Data["port"] = ""
		this.Data["email"] = ""
		this.Data["password"] = ""
		this.Data["isRunning"] = false
	} else {
		this.Data["host"] = rsp.Data.Info.Host
		this.Data["port"] = rsp.Data.Info.Port
		this.Data["email"] = rsp.Data.Info.Username
		this.Data["password"] = rsp.Data.Info.Password
		this.Data["isRunning"] = false
	}

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	Host     string `json:"host"`
	Email    string `json:"email"`
	Port     uint   `json:"port"`
	Password string `json:"password"`
	Must     *actions.Must
}) {
	params.Must.
		Field("email", params.Email).
		Require("请输入邮箱")
	if params.Host == "" {
		this.Fail("请配置邮件服务器")
	}
	if params.Port == 0 {
		this.Fail("请配置邮件端口")
	}
	if params.Password == "" {
		this.Fail("请配置密码")
	}

	res, err := email.SetEmail(&email.SetEmailReq{
		User: &request.UserReq{
			AdminUserId: uint64(this.AdminId()),
		},
		Host:     params.Host,
		Username: params.Email,
		Port:     params.Port,
		Password: params.Password,
	})
	if err != nil || res == nil {
		this.Fail("保存失败")
	}
	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "保存审计邮箱配置 %v", params.Email)

	this.Success()
	// this.Show()
}
