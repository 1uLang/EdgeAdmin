package virus

import (
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/clamav"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"regexp"
	"strings"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	NodeId uint64
}) {
	node, _, err := opnsense_server.GetOpnsenseNodeList()
	if err != nil || node == nil {
		node = make([]*subassemblynode.Subassemblynode, 0)
		//this.ErrorPage(err)
		//return
	}
	//nat 规则列表
	if params.NodeId == 0 && len(node) > 0 {
		params.NodeId = node[0].Id
	}
	version, err := clamav.GetClamAV(&clamav.NodeReq{
		NodeId: params.NodeId,
	})
	if err != nil || version == nil {
		//this.Show()
		//this.ErrorPage(err)
		//return
		this.Data["version"] = maps.Map{
			"update_time":  "每天",
			"version":      "版本",
			"all_total":    "总特征数",
			"update_total": "更新特征数",
			"main_total":   "主要特征数",
		}
	} else {
		//this.Data["version"] = version.Version
		on := version.Version.Main
		re, _ := regexp.Compile("\\d{1,}")
		dataSlice := re.FindAll([]byte(on), -1)
		all_total := "0"
		if len(dataSlice) >= 2 {
			all_total = string(dataSlice[1])
		}

		dataSlice = re.FindAll([]byte(version.Version.Daily), -1)
		update_total := "0"
		if len(dataSlice) >= 2 {
			update_total = string(dataSlice[1])
		}

		day := strings.Split(version.Version.Main, " on ")
		dayly := ""
		if len(day) >= 2 {
			dayly = day[1]
		}
		this.Data["version"] = maps.Map{
			"update_time":  dayly,
			"version":      version.Version.Clamav,
			"all_total":    version.Version.Signatures,
			"update_total": update_total,
			"main_total":   all_total,
		}
	}
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId

	this.Show()
	//this.Success()
}
