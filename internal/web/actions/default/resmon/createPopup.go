package resmon

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{}) {

	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	DomainNames     []string
	Protocols       []string
	CertIdsJSON     []byte
	OriginsJSON     []byte
	RequestHostType int32
	RequestHost     string
	CacheCondsJSON  []byte

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {

	this.Success()
}
