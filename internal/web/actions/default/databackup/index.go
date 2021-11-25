package databackup

import (
	"bytes"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/oplogs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Dirpath string // 为空时表示根目录
}) {

	this.Data["list"] = []int64{}
	this.Data["quota"] = ""
	this.Data["used"] = ""
	this.Data["percent"] = ""
	this.Data["title"] = []int64{}

	defer this.Show()
	// 获取token
	token, err := model.QueryTokenByUID(this.AdminId(), 1)
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}

	// 文件列表（不包含目录）
	// list, err := request.ListFolders(token)
	list, err := request.ListFoldersWithPath(token, params.Dirpath)
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	this.Data["list"] = list.List
	this.Data["quota"] = list.Quota
	this.Data["used"] = list.Used
	this.Data["percent"] = list.Percent
	this.Data["title"] = list.DirList

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	UploadFile *actions.File `json:"uploadFile"`
	Dirpath    string
}) {
	// 获取token
	token, err := model.QueryTokenByUID(this.AdminId(), 1)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 上传文件
	if params.UploadFile == nil {
		this.Fail("请选择要上传的文件")
	}
	upFile, err := params.UploadFile.Read()
	if err != nil {
		this.Fail("读取文件内容错误，请重新上传")
	}
	name := params.UploadFile.Filename
	// err = request.UploadFile(token, name, bytes.NewReader(upFile))
	err = request.UploadFileWithPath(token, name, bytes.NewReader(upFile), params.Dirpath)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "上传backup文件 %v", name)

	this.Success()
	// this.Show()
}
