package web

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/webutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("web")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	webConfig, err := webutils.FindWebConfigWithServerId(this.Parent(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webConfig"] = webConfig

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId int64
	WebId    int64
	Root     string

	Must *actions.Must
}) {

	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWeb(this.AdminContext(), &pb.UpdateHTTPWebRequest{
		WebId: params.WebId,
		Root:  params.Root,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}