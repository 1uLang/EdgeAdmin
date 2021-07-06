package host

import (
	"fmt"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
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
	HostIds []uint64
	Must    *actions.Must
}) {
	params.Must.
		Field("hostIds", params.HostIds).
		Require("请选择高防IP")

	err := host_status_server.DeleteAddr(params.HostIds)
	if err != nil {
		this.Error(fmt.Sprintf("添加主机失败：%v", err), 400)
		return
	}
	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "删除高防IP %v", params.HostIds)
	this.Success()
}
