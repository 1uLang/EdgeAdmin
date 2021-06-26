package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	examine_server "github.com/1uLang/zhiannet-api/hids/server/examine"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"strings"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageNo       int
	PageSize     int
	UserName     string
	State        int
	Score        int
	Type         int
	StartTime    string //体检开始时间
	EndTime      string //体检结束时间
	ExamineItems string //体检项目集合

}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
	}
	req := &examine.SearchReq{}
	req.UserName = "luobing"
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize

	req.State = params.State - 1
	req.Score = params.Score - 1
	if params.Type == 0 {
		req.Type = -1
	} else {
		req.Type = params.Type
	}

	//req.StartTime = params.StartTime
	//req.EndTime = params.EndTime
	//去掉 ','
	params.ExamineItems = strings.TrimPrefix(params.ExamineItems, ",")
	params.ExamineItems = strings.TrimSuffix(params.ExamineItems, ",")
	if len(params.ExamineItems) > 0 {
		req.ExamineItems = strings.Split(params.ExamineItems, ",")
	}
	list, err := examine_server.List(req)
	if err != nil {
		this.ErrorPage(err)
	}
	for k, v := range list.ServerExamineResultInfoList {
		os, err := server.Info(v["serverExamineResultInfo"].(map[string]interface{})["serverIp"].(string))
		if err != nil {
			this.ErrorPage(err)
		}
		list.ServerExamineResultInfoList[k]["os"] = os
	}
	this.Data["datas"] = list.ServerExamineResultInfoList
	this.Data["state"] = params.State
	this.Data["Type"] = params.Type
	this.Data["score"] = params.Score
	this.Data["examineItems"] = params.ExamineItems
	this.Show()
}
