package ddos

import (
	"github.com/1uLang/zhiannet-api/common/server/subassemblynode"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Keyword string
}) {
	list, err := subassemblynode.GetNodeList()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["list"] = list
	//this.Show()
	this.Success()
}
