package invade

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"sync"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) RunGet(params struct{}) {

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
	}
	dashboard := make([]map[string]interface{}, 8)

	invadeLock := sync.RWMutex{}
	invadeWg := sync.WaitGroup{}
	nameUrls := []map[string]string{
		{"name": "病毒木马", "url": "virus"},
		{"name": "网页后门", "url": "webShell"},
		{"name": "反弹shell", "url": "reboundShell"},
		{"name": "异常账号", "url": "abnormalAccount"},
		{"name": "日志异常删除", "url": "logDelete"},
		{"name": "异常登录", "url": "abnormalLogin"},
		{"name": "异常进程", "url": "abnormalProcess"},
		{"name": "系统命令篡改", "url": "systemCmd"},
	}
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
	args.UserName = "cysct56"
	args.PageSize = 1

	for i, f := range fns {
		invadeWg.Add(1)
		go func(idx int, fn func(*risk.RiskSearchReq) (risk.RiskSearchResp, error)) {
			defer invadeWg.Done()
			risk, _ := fn(args)
			invadeLock.Lock()
			dashboard[idx] = map[string]interface{}{
				"name":  nameUrls[idx]["name"],
				"count": risk.TotalData,
				"url":   nameUrls[idx]["url"],
			}
			invadeLock.Unlock()
			fmt.Println(idx, nameUrls[idx]["name"], "=============", risk)
		}(i, f)
	}
	invadeWg.Wait()
	this.Data["dashboard"] = dashboard
	this.Show()
}
