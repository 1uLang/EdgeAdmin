package whiteblacklist

import (
	"fmt"
	black_white_list_server "github.com/1uLang/zhiannet-api/ddos/server/black_white_list"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Address  string
	NodeId   uint64
	Page     int64
	PageSize int64
}) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 20
	}
	defer this.Show()
	this.Data["list"] = nil
	this.Data["ddos"] = nil
	this.Data["Address"] = ""
	this.Data["nodeId"] = ""

	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	if len(ddos) == 0 {
		this.Data["errorMessage"] = "未配置DDoS防火墙节点"
		return
	}
	this.Data["ddos"] = ddos
	if params.NodeId == 0 {
		params.NodeId = ddos[0].Id
	}
	//TODO ddos黑名单查询支持分页查询，但是每页固定500条，等保平台需要自定义每页条数翻页查询
	pageNum := int64(0)
	if ((params.Page*params.PageSize)/500) > 0 && ((params.Page*params.PageSize)%500) > 0 {
		pageNum = ((params.Page * params.PageSize) / 500)
	}
	if ((params.Page*params.PageSize)/500) > 0 && ((params.Page*params.PageSize)%500) == 0 {
		pageNum = ((params.Page * params.PageSize) / 500) - 1
		if pageNum < 0 {
			pageNum = 0
		}
	}
	req := &black_white_list_server.BWReq{
		NodeId: params.NodeId,
		Addr:   params.Address,
		Page:   int(pageNum),
	}

	this.Data["nodeId"] = req.NodeId
	list, err := black_white_list_server.GetBWList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取DDoS防火墙黑白名单列表失败：%v", err.Error())
		return
	}
	page := this.NewPage(int64(len(list.Bwlist) + int(pageNum*500)))
	if len(list.Bwlist) == 500 {
		page = this.NewPage(int64(len(list.Bwlist) + int(pageNum*500) + 1))
	}
	this.Data["page"] = page.AsHTML()

	offset := page.Offset
	offset -= (pageNum * 500)
	end := offset + page.Size
	fmt.Println("page.Offset", page.Offset, "page.Size", page.Size)
	if end > page.Total {
		end = page.Total
	}
	this.Data["list"] = list.Bwlist[offset:end]
	this.Data["Address"] = list.Address

}
