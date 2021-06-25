package invade

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//入侵威胁

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeHids)).
			Data("teaMenu", "hids").
			Prefix("/hids/invade").
			Get("", new(IndexAction)).
			//病毒木马
			GetPost("/virus", new(VirusAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/virus/detailList", new(VirusDetailListAction)). //病毒木马 详情列表
			Get("/virus/detail", new(VirusDetailAction)).         //病毒木马 详情
			//网页后门
			GetPost("/webShell", new(WebShellAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/webShell/detailList", new(WebShellDetailListAction)). //网页后门 详情列表
			Get("/webShell/detail", new(WebShellDetailAction)).         //网页后门 详情
			//反弹shell
			GetPost("/reboundShell", new(ReboundShellAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/reboundShell/detailList", new(ReboundShellDetailListAction)). //反弹shell 详情列表
			Get("/reboundShell/detail", new(ReboundShellDetailAction)).         //反弹shell 详情
			//异常账号
			GetPost("/abnormalAccount", new(AbnormalAccountAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/abnormalAccount/detailList", new(AbnormalAccountDetailListAction)). //异常账号 详情列表
			Get("/abnormalAccount/detail", new(AbnormalAccountDetailAction)).         //异常账号 详情
			//日志异常删除
			GetPost("/logDelete", new(LogDeleteAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/logDelete/detailList", new(LogDeleteDetailListAction)). //日志异常删除 详情列表
			Get("/logDelete/detail", new(LogDeleteDetailAction)).         //日志异常删除 详情
			//异常登录
			GetPost("/abnormalLogin", new(AbnormalLoginAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/abnormalLogin/detailList", new(AbnormalLoginDetailListAction)). //异常登录 详情列表
			Get("/abnormalLogin/detail", new(AbnormalLoginDetailAction)).         //异常登录 详情
			//异常进程
			GetPost("/abnormalProcess", new(AbnormalProcessAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/abnormalProcess/detailList", new(AbnormalProcessDetailListAction)). //异常进程 详情列表
			Get("/abnormalProcess/detail", new(AbnormalProcessDetailAction)).         //异常进程 详情
			//系统命令篡改
			GetPost("/logDelete", new(LogDeleteAction)).                  //相关主机 - 操作[信任-隔离]
			Get("/logDelete/detailList", new(LogDeleteDetailListAction)). //系统命令篡改 详情列表
			Get("/logDelete/detail", new(LogDeleteDetailAction)).         //系统命令篡改 详情
			EndAll()
	})
}
