package expire

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/index"
	"github.com/iwind/TeaGo/maps"
	"strconv"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "expire", "expire")
}

func (this *IndexAction) RunGet(params struct {
	NodeId int64
}) {

	conf, err := index.LoadServerExpireConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var expireTime int64
	if conf.Expire.On && conf.Expire.Time != "" {
		expireTime, _ = strconv.ParseInt(conf.Expire.Time, 10, 64)
	}
	this.Data["info"] = maps.Map{
		"dayTo":     time.Unix(expireTime, 0).Format("2006-01-02"),
		"name":      "CloudSafe",
		"code":      conf.Expire.Code,
		"package":   conf.Expire.Num,
		"isExpired": expireTime < time.Now().Unix(),
		"isOn":      conf.Expire.On,
	}

	this.Show()
}
