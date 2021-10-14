package feature

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	NodeId uint64
}) {
	this.Data["version"] = maps.Map{
		"update_time": time.Now().Format("2006-01-02"),
	}

	this.Show()

}
