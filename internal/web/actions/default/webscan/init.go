package webscan

import (
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/awvs/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//web漏洞扫描
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeWebScan)).
			Data("teaMenu", "webscan").
			Prefix("/webscan").
			Get("", new(IndexAction)).
			EndAll()
	})
	info, err := server.GetWebScan()
	if err != nil {
		panic("漏扫节点获取失败")
	}
	Key = info.Key
	ServerUrl = info.Addr
}

var ServerUrl = "" //"https://scan-web.zhiannet.com"
var Key = ""

func InitAPIServer() error {

	err := server.SetUrl(ServerUrl)
	if err != nil {
		return err
	}
	//初始化 awvs 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		XAuth: Key,
	})
	if err != nil {
		return err
	}
	return nil
}
