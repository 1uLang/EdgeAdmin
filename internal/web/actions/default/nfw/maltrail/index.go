package maltrail

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/maltrail/server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Date     string
	Page     int
	PageSize int
}) {
	if params.Date == "" {
		params.Date = time.Now().Format("2006-01-02")
	}
	list := make([]*server.ListResp, 0)
	key := fmt.Sprintf("maltrail-list-%v", params.Date)
	lists, err := cache.CheckCache(key, func() (interface{}, error) {
		list, err := server.GetList(&server.ListReq{
			Date: params.Date,
		})
		return list, err
	}, 60, true)

	b, err := json.Marshal(lists)
	if err == nil {
		json.Unmarshal(b, &list)
	}

	page := this.NewPage(int64(len(list)))
	this.Data["page"] = page.AsHTML()

	this.Data["date"] = params.Date
	if params.PageSize == 0 {
		params.PageSize = 20
	}
	if params.Page <= 1 {
		params.Page = 1
	}
	if len(list) > params.PageSize {
		startKey := ((params.Page - 1) * params.PageSize)
		if startKey > len(list)-1 {
			this.Data["list"] = list
			this.Show()
		}
		endKey := startKey + params.PageSize
		if endKey > len(list)-1 {
			endKey = len(list) - 1
		}
		list = list[startKey:endKey]
	}
	this.Data["list"] = list
	this.Show()

}
