package nfw

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/global_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"strconv"
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
	status, err := global_status.GetGlobalStatus(&global_status.GlobalReq{
		NodeId: params.NodeId,
	})
	if err != nil || status == nil {
		//this.Show()
		//this.ErrorPage(err)
		//return
		status = &global_status.GlobalResp{}
	}
	global := map[string]interface{}{}
	global["verInfo"] = strings.Join(status.Versions, " ") //版本信息
	global["cpuInfo"] = status.CPU.Model                   //cpu byte
	global["slbInfo"] = strings.Join(status.CPU.Load, ",") //cpu 负载
	s, _ := strconv.ParseInt(status.Uptime, 10, 64)
	global["runTime"] = toDay(s)                            //运行时间
	global["timeNow"] = status.DateFrmt                     //当前时间
	global["lastConfigTime"] = status.Config.LastChangeFrmt //最后一次配置时间
	global["cpuPer"] = status.CPU.Used                      //cpu 使用率
	global["statusCount"] = map[string]string{"maxValue": status.Kernel.Pf.Maxstates, "useValue": status.Kernel.Pf.States}
	global["MBUFCount"] = map[string]string{"maxValue": status.Kernel.Mbuf.Max, "useValue": status.Kernel.Mbuf.Total}
	MemoryTotal, _ := strconv.ParseInt(status.Kernel.Memory.Total, 10, 64)
	MemoryTotal = MemoryTotal / 1024 / 1024
	MemoryUsed, _ := strconv.ParseInt(status.Kernel.Memory.Used, 10, 64)
	MemoryUsed = MemoryUsed / 1024 / 1024
	global["RAMPer"] = map[string]string{"maxValue": fmt.Sprintf("%v", MemoryTotal), "useValue": fmt.Sprintf("%v", MemoryUsed), "uint": "MB"}
	var swapTotal int64
	var swapUse int64
	if len(status.Disk.Swap) > 0 {
		swapTotal, _ = strconv.ParseInt(status.Disk.Swap[0].Total, 10, 64)
		swapTotal = swapTotal / 1024 / 1024
		swapUse, _ = strconv.ParseInt(status.Disk.Swap[0].Used, 10, 64)
		swapUse = swapUse / 1024 / 1024
	}
	global["SWAPPer"] = map[string]string{"maxValue": fmt.Sprintf("%v", swapTotal), "useValue": fmt.Sprintf("%v", swapUse), "uint": "MB"}
	diskPer := "0"
	if len(status.Disk.Devices) > 0 {
		diskPer = strings.Replace(status.Disk.Devices[0].Capacity, "%", "", -1)
	}
	global["diskPer"] = diskPer

	this.Data["tableData"] = global
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	this.Show()
	//this.Success()
}

//秒转运行时间
func toDay(s int64) string {
	var day string
	day = fmt.Sprintf("%v天 %v小时 %v分", (s / (60 * 60 * 24)), (s%(60*60*24))/(60*60), (s%(60*60))/(60))

	return day
}
