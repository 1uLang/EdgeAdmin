package whiteblacklist

import (
	black_white_list_server "github.com/1uLang/zhiannet-api/ddos/server/black_white_list"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DelAction struct {
	actionutils.ParentAction
}

func (this *DelAction) Init() {
	this.Nav("", "", "")
}

func (this *DelAction) RunGet(params struct{}) {
	this.Show()
}

func (this *DelAction) RunPost(params struct {
	Addr   string
	NodeId uint64
	Type   string
	Must   *actions.Must
}) {
	params.Must.
		Field("Addr", params.Addr).
		Require("请输入ip地址")

	_, err := black_white_list_server.DeleteBW(&black_white_list_server.EditBWReq{
		Addr:   []string{params.Addr},
		NodeId: params.NodeId,
		White:  !(params.Type == "blacklist"),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
