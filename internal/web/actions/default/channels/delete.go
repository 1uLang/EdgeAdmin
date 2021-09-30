package channels

import (
	"github.com/1uLang/zhiannet-api/common/server/channels_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_users_server"
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
	Id uint64 `json:"id"`

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("请输入id")
	if params.Id <= 0 {
		this.Fail("id必须大于0")
	}
	total, err := edge_users_server.GetChannelUserTotal([]uint64{params.Id}, false)
	if err != nil {
		this.Fail("获取渠道用户数失败")
	}
	if num, ok := total[params.Id]; ok && num > 0 {
		this.Fail("渠道已有关联用户，无法删除")
	}

	err = channels_server.Del(params.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
