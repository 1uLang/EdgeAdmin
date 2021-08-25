package resmon

import (
	"github.com/1uLang/zhiannet-api/resmon/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) Init() {
	this.Nav("", "", "")
}

func (this *DeleteAction) RunPost(params struct {
	Id   string
	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Id).
		Require("请输入主机Id")

	err := request.DeleteAgent(params.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
