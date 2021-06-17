package vulnerabilities

import (
	"github.com/1uLang/zhiannet-api/awvs/model/vulnerabilities"
	vulnerabilities_server "github.com/1uLang/zhiannet-api/awvs/server/vulnerabilities"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/webscan"
)

//任务目标
type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	pageSize int
	pageNo   int
	address  string
	level    string
}) {
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.pageNo < 0 {
		params.pageNo = 0
	}
	if params.pageSize < 0 {
		params.pageSize = 20
	}
	var query string
	if params.address != "" {
		query += "target_id:" + params.address
		query += ";"
	}
	if params.level != "" {
		query += "severity:" + params.level
		query += ";"
	}

	list, err := vulnerabilities_server.List(&vulnerabilities.ListReq{Limit: params.pageSize, C: params.pageNo * params.pageSize, Query: query})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["vulnerabilities"] = list["vulnerabilities"]
	this.Show()
}
