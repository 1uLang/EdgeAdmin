package command

import "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "hids", "index")
}

func (this *IndexAction) RunGet() {
	//err := InitAPIServer()
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	this.Show()
}
