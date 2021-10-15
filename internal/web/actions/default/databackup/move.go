package databackup

import (
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type RenameAction struct {
	actionutils.ParentAction
}

func (this *RenameAction) Init() {
	this.Nav("", "", "")
}

func (this *RenameAction) RunGet(params struct {
	Name string
	Url  string
}) {
	this.Data["name"] = params.Name
	this.Data["url"] = params.Url
	this.Show()
}

func (this *RenameAction) RunPost(params struct {
	SrcURL  string
	NewName string

	Must *actions.Must
}) {
	params.Must.
		Field("srcURL", params.SrcURL).
		Require("请输入原文件或文件夹的url路径").
		Field("newName", params.NewName).
		Require("请输入新的的名字")
	if params.SrcURL == "" || params.NewName == "" {
		this.Fail("url路径，新名字不能为空")
	}
	// 获取token
	token, err := model.QueryTokenByUID(this.AdminId(), 1)
	if err != nil {
		this.FailField("error", err.Error())
		return
	}

	err = request.MoveFileOrFolder(params.SrcURL, params.NewName, token)
	if err != nil {
		this.FailField("error", err.Error())
		return
	}

	this.Success()
}
