package monitor

import "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	DomainId int64
}) {
	// 记录日志
	//defer this.CreateLog(oplogs.LevelInfo, "从DNS服务商中删除域名 %d", params.DomainId)

	this.Success()
}
