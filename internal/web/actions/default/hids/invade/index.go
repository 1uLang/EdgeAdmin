package invade

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) RunGet(params struct{}) {

	defer this.Show()

	dashboard := []map[string]interface{}{
		{"name": "病毒木马", "url": "virus", "count": 0},
		{"name": "网页后门", "url": "webShell", "count": 0},
		{"name": "反弹shell", "url": "reboundShell", "count": 0},
		{"name": "异常账号", "url": "abnormalAccount", "count": 0},
		{"name": "日志异常删除", "url": "logDelete", "count": 0},
		{"name": "异常登录", "url": "abnormalLogin", "count": 0},
		{"name": "异常进程", "url": "abnormalProcess", "count": 0},
		{"name": "系统命令篡改", "url": "systemCmd", "count": 0},
	}

	this.Data["dashboard"] = dashboard

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}

	//invadeLock := sync.RWMutex{}
	//invadeWg := sync.WaitGroup{}

	fns := []func(*risk.RiskSearchReq) (risk.RiskSearchResp, error){
		risk.VirusList,
		risk.WebShellList,
		risk.ReboundList,
		risk.AbnormalAccountList,
		risk.LogDeleteList,
		risk.AbnormalLoginList,
		risk.AbnormalProcessList,
		risk.SystemCmdList,
	}
	args := &risk.RiskSearchReq{}

	args.UserName, err = this.UserName()
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取用户信息失败：%v", err)
		return
	}
	args.PageSize = 1

	for idx, fn := range fns {
		//invadeWg.Add(1)
		//go func(idx int, fn func(*risk.RiskSearchReq) (risk.RiskSearchResp, error)) {
		//	defer invadeWg.Done()
			risk, _ := fn(args)
			//invadeLock.Lock()
			dashboard[idx]["count"] = risk.TotalData
			//invadeLock.Unlock()
		//}(i, f)
	}
	//invadeWg.Wait()

	this.Data["dashboard"] = dashboard

}
