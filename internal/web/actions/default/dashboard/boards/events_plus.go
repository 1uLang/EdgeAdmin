// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package boards

import (
	teaconst "github.com/TeaOSLab/EdgeAdmin/internal/const"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/dashboard/boards/boardutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/nodeconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type EventsAction struct {
	actionutils.ParentAction
}

func (this *EventsAction) Init() {
	this.Nav("", "", "event")
}

func (this *EventsAction) RunGet(params struct{}) {
	if !teaconst.IsPlus {
		this.RedirectURL("/dashboard")
		return
	}

	this.Data["keyword"] = ""

	// 初始化
	err := boardutils.InitBoard(this.Parent())
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 分页
	countResp, err := this.RPC().NodeLogRPC().CountNodeLogs(this.AdminContext(), &pb.CountNodeLogsRequest{
		IsUnread: true,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var count = countResp.Count
	var page = this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	// 单页
	logsResp, err := this.RPC().NodeLogRPC().ListNodeLogs(this.AdminContext(), &pb.ListNodeLogsRequest{
		IsUnread: true,
		Offset:   page.Offset,
		Size:     page.Size,
	})
	if err != nil {

	}
	var logMaps = []maps.Map{}
	for _, log := range logsResp.NodeLogs {
		var logMap = maps.Map{
			"id":          log.Id,
			"role":        log.Role,
			"tag":         log.Tag,
			"description": log.Description,
			"createdTime": timeutil.FormatTime("Y-m-d H:i:s", log.CreatedAt),
			"level":       log.Level,
			"isToday":     timeutil.FormatTime("Y-m-d", log.CreatedAt) == timeutil.Format("Y-m-d"),
			"count":       log.Count,
		}
		switch log.Role {
		case nodeconfigs.NodeRoleNode:
			// 节点信息
			nodeResp, err := this.RPC().NodeRPC().FindEnabledNode(this.AdminContext(), &pb.FindEnabledNodeRequest{NodeId: log.NodeId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			node := nodeResp.Node
			if node == nil || node.NodeCluster == nil {
				_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{NodeLogIds: []int64{log.Id}})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				continue
			}
			logMap["node"] = maps.Map{
				"id": node.Id,
				"cluster": maps.Map{
					"id":   node.NodeCluster.Id,
					"name": node.NodeCluster.Name,
				},
				"name": node.Name,
			}

			// 服务信息
			var serverMap = maps.Map{"id": 0}
			if log.ServerId > 0 {
				serverResp, err := this.RPC().ServerRPC().FindEnabledUserServerBasic(this.AdminContext(), &pb.FindEnabledUserServerBasicRequest{ServerId: log.ServerId})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				var server = serverResp.Server
				if server != nil {
					serverMap = maps.Map{"id": server.Id, "name": server.Name}
				}
			}
			logMap["server"] = serverMap
		case nodeconfigs.NodeRoleAPI:
			nodeResp, err := this.RPC().APINodeRPC().FindEnabledAPINode(this.AdminContext(), &pb.FindEnabledAPINodeRequest{ApiNodeId: log.NodeId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			var node = nodeResp.ApiNode
			if node == nil {
				_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{NodeLogIds: []int64{log.Id}})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				continue
			}
			logMap["node"] = maps.Map{
				"id":   node.Id,
				"name": node.Name,
			}
		case nodeconfigs.NodeRoleDNS:
			// 节点信息
			nodeResp, err := this.RPC().NSNodeRPC().FindEnabledNSNode(this.AdminContext(), &pb.FindEnabledNSNodeRequest{NsNodeId: log.NodeId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			node := nodeResp.NsNode
			if node == nil || node.NsCluster == nil {
				_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{NodeLogIds: []int64{log.Id}})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				continue
			}
			logMap["node"] = maps.Map{
				"id": node.Id,
				"cluster": maps.Map{
					"id":   node.NsCluster.Id,
					"name": node.NsCluster.Name,
				},
				"name": node.Name,
			}
		case nodeconfigs.NodeRoleReport:
			nodeResp, err := this.RPC().ReportNodeRPC().FindEnabledReportNode(this.AdminContext(), &pb.FindEnabledReportNodeRequest{ReportNodeId: log.NodeId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			var node = nodeResp.ReportNode
			if node == nil {
				_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{NodeLogIds: []int64{log.Id}})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				continue
			}
			logMap["node"] = maps.Map{
				"id":   node.Id,
				"name": node.Name,
			}
		case nodeconfigs.NodeRoleMonitor:
			nodeResp, err := this.RPC().MonitorNodeRPC().FindEnabledMonitorNode(this.AdminContext(), &pb.FindEnabledMonitorNodeRequest{MonitorNodeId: log.NodeId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			var node = nodeResp.MonitorNode
			if node == nil {
				_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{NodeLogIds: []int64{log.Id}})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				continue
			}
			logMap["node"] = maps.Map{
				"id":   node.Id,
				"name": node.Name,
			}
		case nodeconfigs.NodeRoleUser:
			nodeResp, err := this.RPC().UserNodeRPC().FindEnabledUserNode(this.AdminContext(), &pb.FindEnabledUserNodeRequest{UserNodeId: log.NodeId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			var node = nodeResp.UserNode
			if node == nil {
				_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{NodeLogIds: []int64{log.Id}})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				continue
			}
			logMap["node"] = maps.Map{
				"id":   node.Id,
				"name": node.Name,
			}
		case nodeconfigs.NodeRoleAdmin:
			// TODO
		case nodeconfigs.NodeRoleDatabase:
			nodeResp, err := this.RPC().DBNodeRPC().FindEnabledDBNode(this.AdminContext(), &pb.FindEnabledDBNodeRequest{DbNodeId: log.NodeId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			var node = nodeResp.DbNode
			if node == nil {
				_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{NodeLogIds: []int64{log.Id}})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				continue
			}
			logMap["node"] = maps.Map{
				"id":   node.Id,
				"name": node.Name,
			}
		default:
			_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{NodeLogIds: []int64{log.Id}})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			continue
		}

		logMaps = append(logMaps, logMap)
	}
	this.Data["logs"] = logMaps

	this.Show()
}
