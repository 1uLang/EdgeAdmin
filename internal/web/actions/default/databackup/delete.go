package databackup

import (
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) Init() {
	this.Nav("", "", "")
}

func (this *DeleteAction) RunPost(params struct {
	Fp string
}) {
	// 获取token
	token, err := model.QueryTokenByUID(this.AdminId(), 1)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 删除文件
	// err = request.DeleteFile(token, params.Name)
	err = request.DeleteFileWithPath(token, params.Fp)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	defer this.CreateLogInfo("删除数据备份文件 %v", params.Fp)

	this.Success()
}
