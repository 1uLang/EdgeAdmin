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
	PageSize int
	PageNo   int
	Address  string
	Severity string

	List bool
}) {
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 20
	}
	var query string
	if params.Address != "" {
		query += "target_id:" + params.Address
		query += ";"
	}
	if params.Severity != "" {
		query += "severity:" + params.Severity
		query += ";"
	}

	list, err := vulnerabilities_server.List(&vulnerabilities.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, Query: query, AdminUserId: uint64(this.AdminId())})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//this.Data["vulnerabilities"] = list["vulnerabilities"]
	if lists, ok := list["vulnerabilities"]; ok {
		this.Data["vulnerabilities"] = lists
	}
	if !params.List {
		this.Show()
	} else {
		this.Success()
	}
}
