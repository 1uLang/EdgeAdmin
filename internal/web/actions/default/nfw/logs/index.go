package logs

import (
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	req_logs "github.com/1uLang/zhiannet-api/opnsense/request/logs"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/logs"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Keyword  string
	NodeId   uint64
	PageNum  int
	PageSize int
}) {
	node, _, err := opnsense_server.GetOpnsenseNodeList()
	if err != nil || node == nil {
		//this.ErrorPage(err)
		//return
		node = make([]*subassemblynode.Subassemblynode, 0)
	}
	// 规则列表
	if params.NodeId == 0 && len(node) > 0 {
		params.NodeId = node[0].Id
	}
	list, err := logs.GetLogsList(&logs.LogReq{
		NodeId:   params.NodeId,
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
		Keyword:  params.Keyword,
	})
	if err != nil || list == nil {
		//this.ErrorPage(err)
		//return
		list = &req_logs.LogListResp{}
	}
	if len(list.Rows) > 0 {
		this.Data["tableData"] = list.Rows
	} else {
		this.Data["tableData"] = make([]interface{}, 0)
	}
	this.Data["total"] = list.Total  //总条数
	this.Data["num"] = list.RowCount //当前页条数
	this.Data["keyword"] = params.Keyword
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	this.Show()
}
