package databackup

import (
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DirAction struct {
	actionutils.ParentAction
}

func (this *DirAction) Init() {
	this.Nav("", "", "")
}

func (this *DirAction) RunPost(params struct {
	Purl string
	Name string
}) {
	// 获取token
	token, err := model.QueryTokenByUID(this.AdminId(), 1)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	err = request.CreateFoler(token,params.Purl,params.Name)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
