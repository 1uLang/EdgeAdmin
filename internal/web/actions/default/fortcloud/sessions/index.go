package sessions

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

	this.Data["online"] = []interface{}{}
	defer this.Show()
	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.Data["nodeErr"] = err.Error()
		return
	}
	online, err := req.Session.List(&session_model.ListReq{
		Status:      "connected",
		AdminUserId: uint64(this.AdminId()),
		PageSize:    999,
		PageIndex:   1,
	})
	if err != nil {
		this.Data["nodeErr"] = err.Error()
		return
	}
	this.Data["online"] = online

}
