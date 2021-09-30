package channels

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"github.com/1uLang/zhiannet-api/common/server/channels_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_users_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "channels", "index")
}

func (this *IndexAction) RunGet(params struct {
	Keyword  string
	PageNum  int
	PageSize int
	Status   string
}) {
	if params.PageSize == 0 {
		params.PageSize = 20
	}
	if params.PageNum == 0 {
		params.PageNum = 1
	}
	list, count, err := channels_server.GetList(&channels.ChannelReq{
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
		Status:   params.Status,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	listMap := make([]interface{}, len(list))
	if len(list) > 0 {
		chanIds := make([]uint64, 0)
		for _, v := range list {
			chanIds = append(chanIds, v.Id)
		}
		total, err := edge_users_server.GetChannelUserTotal(chanIds, false)
		if err != nil {
			this.ErrorPage(fmt.Errorf("获取渠道用户数失败"))
		}
		for k, v := range list {
			clientCount := int64(0)
			if num, ok := total[v.Id]; ok {
				clientCount = num
			}
			createTime := time.Unix(int64(v.CreateTime), 0).Format("2006-01-02 15:04:05")
			listMap[k] = map[string]interface{}{
				"id":          v.Id,
				"name":        v.Name,
				"user":        v.User,
				"mobile":      v.Mobile,
				"productName": v.ProductName,
				"status":      v.Status,
				"createTime":  createTime,
				"domain":      v.Domain,
				"logo":        v.Logo,
				"remake":      v.Remake,
				"clientCount": clientCount,
			}
		}
	}
	this.Data["tableData"] = listMap
	page := this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	Keyword  string
	PageNum  int
	PageSize int
	Status   string
}) {
	if params.PageSize == 0 {
		params.PageSize = 20
	}
	if params.PageNum == 0 {
		params.PageNum = 1
	}
	list, _, err := channels_server.GetList(&channels.ChannelReq{
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
		Status:   params.Status,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	listMap := make([]interface{}, len(list))
	if len(list) > 0 {
		chanIds := make([]uint64, 0)
		for _, v := range list {
			chanIds = append(chanIds, v.Id)
		}
		total, err := edge_users_server.GetChannelUserTotal(chanIds, false)
		if err != nil {
			this.ErrorPage(fmt.Errorf("获取渠道用户数失败"))
		}
		for k, v := range list {
			clientCount := int64(0)
			if num, ok := total[v.Id]; ok {
				clientCount = num
			}
			createTime := time.Unix(int64(v.CreateTime), 0).Format("2006-01-02 15:04:05")
			listMap[k] = map[string]interface{}{
				"id":          v.Id,
				"name":        v.Name,
				"user":        v.User,
				"mobile":      v.Mobile,
				"productName": v.ProductName,
				"status":      v.Status,
				"createTime":  createTime,
				"domain":      v.Domain,
				"logo":        v.Logo,
				"remake":      v.Remake,
				"clientCount": clientCount,
			}
		}
	}
	this.Data["tableData"] = listMap

	this.Success()
}
