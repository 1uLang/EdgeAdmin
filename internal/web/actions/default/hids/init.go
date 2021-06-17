package hids

import (
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//主机防护

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeHids)).
			Data("teaMenu", "hids").
			Prefix("/hids").
			Get("", new(IndexAction)).
			EndAll()
	})
}
func InitAPIServer() error {
	err := server.SetUrl("https://user.cloudhids.net")
	if err != nil {
		return err
	}
	//初始化 awvs 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		AppId:  "39rkz",
		Secret: "tkvgpvjuht2625mo",
	})
	if err != nil {
		return err
	}
	return nil
}
