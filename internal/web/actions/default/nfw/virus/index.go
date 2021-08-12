package virus

import (
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/clamav"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
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
			"daily":      "每天",
			"main":       "主要",
			"bytecode":   "字节码",
			"signatures": "签名总数",
			"clamav":     "引擎版本",
		}
	} else {
		this.Data["version"] = version.Version
	}
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	this.Show()
	//this.Success()
}
