package examine

import (
	"fmt"
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
		return
	}
	req := &examine.SearchReq{}
	req.UserName, err = this.UserName()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取当前用户信息失败：%v", err))
		return
	}
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
		return
	}
	datas := make([]map[string]interface{}, 0)
	for k, v := range list.ServerExamineResultInfoList {
		if v["userName"] != req.UserName {
			continue
		}
		os, err := server.Info(v["serverExamineResultInfo"].(map[string]interface{})["serverIp"].(string), req.UserName)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		list.ServerExamineResultInfoList[k]["os"] = os
		datas = append(datas, list.ServerExamineResultInfoList[k])
	}
	this.Data["datas"] = datas
	this.Data["state"] = params.State
	this.Data["Type"] = params.Type
	this.Data["score"] = params.Score
	this.Data["examineItems"] = params.ExamineItems
	this.Show()
}
