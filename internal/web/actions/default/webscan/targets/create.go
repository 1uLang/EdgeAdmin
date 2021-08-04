package targets

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"

	"github.com/1uLang/zhiannet-api/nessus/model/scans"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}
func (this *CreateAction) RunPost(params struct {
	Address string
	Desc    string
	Type 	int

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("address", params.Address).
		Require("请输入目标地址").
		Match("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", "请输入正确的目标地址")

	params.Must.
		Field("type", params.Type).
		Require("请选择类型")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.Type == 1 {
		req := &targets.AddReq{Address: params.Address, AdminUserId: uint64(this.AdminId())}
		req.Description = params.Desc

		_, err = targets_server.Add(req)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		// 日志
		this.CreateLogInfo("WEB漏洞扫描 - 创建任务目标:[%v]成功", params.Address)
	} else if params.Type == 2{

		req := &scans.AddReq{}
		req.AdminUserId = uint64(this.AdminId())
		req.Settings.Name = params.Address
		req.Settings.Text_targets = params.Address
		req.Settings.Description = params.Desc
		//先忽略错误。防止页面重复提交
		go func() {
			err = scans_server.Create(req)
			if err != nil {
				this.CreateLogInfo("主机漏洞扫描 - 创建任务目标:[%v]失败:%v", params.Address,err)
			}else{
				this.CreateLogInfo("主机漏洞扫描 - 创建任务目标:[%v]成功", params.Address)
			}
			// 日志
		}()

	}else {
		this.ErrorPage(fmt.Errorf("漏洞扫描类型参数错误"))
		return
	}
	this.Success()
}
