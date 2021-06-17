package examine

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) RunPost(params struct {
	ProviderId int64
}) {

	this.Success()
}
