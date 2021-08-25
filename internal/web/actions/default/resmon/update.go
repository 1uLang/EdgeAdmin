package resmon

import (
	"github.com/1uLang/zhiannet-api/resmon/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdateAction) RunPost(params struct {
	Id   string
	Name string
	Host string
	Aid  uint8
	On   bool
	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Id).
		Require("请输入主机Id")

	err := request.UpdateAgent(params.Name, params.Host, params.Id, params.On, params.Aid)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
