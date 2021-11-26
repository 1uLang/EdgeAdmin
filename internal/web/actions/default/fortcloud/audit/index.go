package audit

import (
	session_model "github.com/1uLang/zhiannet-api/next-terminal/model/session"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/fortcloud"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}
func (this *IndexAction) RunGet(params struct {
	PageSize int
	PageNo   int
}) {
	defer this.Show()
	this.Data["offline"] = []int{}
	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}

	offline, err := req.Session.List(&session_model.ListReq{
		Status:      "disconnected",
		AdminUserId: uint64(this.AdminId()),
		PageIndex:   1,
		PageSize:    999,
	})
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	this.Data["offline"] = offline

}
