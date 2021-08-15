// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dashboard

import (
	"github.com/1uLang/zhiannet-api/awvs/server/targets"
	"github.com/1uLang/zhiannet-api/common/server/subassemblynode"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeAdmin/internal/const"
	"github.com/TeaOSLab/EdgeAdmin/internal/utils/numberutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/maps"
	"regexp"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	if teaconst.IsPlus {
		this.RedirectURL("/dashboard/boards")
		return
	}

	// 取得用户的权限
	module, ok := configloaders.FindFirstAdminModule(this.AdminId())
	if ok {
		if module != "dashboard" {
			for _, m := range configloaders.AllModuleMaps() {
				if m.GetString("code") == module {
					this.RedirectURL(m.GetString("url"))
					return
				}
			}
		}
	}

	// 读取看板数据
	resp, err := this.RPC().AdminRPC().ComposeAdminDashboard(this.AdminContext(), &pb.ComposeAdminDashboardRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["dashboard"] = maps.Map{
		"countServers":      resp.CountServers,
		"countNodeClusters": resp.CountNodeClusters,
		"countNodes":        resp.CountNodes,
		"countUsers":        resp.CountUsers,
		"countAPINodes":     resp.CountAPINodes,
		"countDBNodes":      resp.CountDBNodes,
		"countUserNodes":    resp.CountUserNodes,

		"canGoServers":  configloaders.AllowModule(this.AdminId(), configloaders.AdminModuleCodeServer),
		"canGoNodes":    configloaders.AllowModule(this.AdminId(), configloaders.AdminModuleCodeServer),
		"canGoSettings": configloaders.AllowModule(this.AdminId(), configloaders.AdminModuleCodeSetting),
		"canGoUsers":    configloaders.AllowModule(this.AdminId(), configloaders.AdminModuleCodeUser),
	}

	// 今日流量
	todayTrafficBytes := int64(0)
	if len(resp.DailyTrafficStats) > 0 {
		todayTrafficBytes = resp.DailyTrafficStats[len(resp.DailyTrafficStats)-1].Bytes
	}
	todayTrafficString := numberutils.FormatBytes(todayTrafficBytes)
	result := regexp.MustCompile(`^(?U)(.+)([a-zA-Z]+)$`).FindStringSubmatch(todayTrafficString)
	if len(result) > 2 {
		this.Data["todayTraffic"] = result[1]
		this.Data["todayTrafficUnit"] = result[2]
	} else {
		this.Data["todayTraffic"] = todayTrafficString
		this.Data["todayTrafficUnit"] = ""
	}

	// 24小时流量趋势
	{
		statMaps := []maps.Map{}
		for _, stat := range resp.HourlyTrafficStats {
			statMaps = append(statMaps, maps.Map{
				"bytes": stat.Bytes,
				"hour":  stat.Hour[8:],
			})
		}
		this.Data["hourlyTrafficStats"] = statMaps
	}

	// 15天流量趋势
	{
		statMaps := []maps.Map{}
		for _, stat := range resp.DailyTrafficStats {
			statMaps = append(statMaps, maps.Map{
				"bytes": stat.Bytes,
				"day":   stat.Day[4:6] + "月" + stat.Day[6:] + "日",
			})
		}
		this.Data["dailyTrafficStats"] = statMaps
	}

	// 版本升级
	if resp.NodeUpgradeInfo != nil {
		this.Data["nodeUpgradeInfo"] = maps.Map{
			"count":   resp.NodeUpgradeInfo.CountNodes,
			"version": resp.NodeUpgradeInfo.NewVersion,
		}
	} else {
		this.Data["nodeUpgradeInfo"] = maps.Map{
			"count":   0,
			"version": "",
		}
	}
	if resp.MonitorNodeUpgradeInfo != nil {
		this.Data["monitorNodeUpgradeInfo"] = maps.Map{
			"count":   resp.MonitorNodeUpgradeInfo.CountNodes,
			"version": resp.MonitorNodeUpgradeInfo.NewVersion,
		}
	} else {
		this.Data["monitorNodeUpgradeInfo"] = maps.Map{
			"count":   0,
			"version": "",
		}
	}
	if resp.ApiNodeUpgradeInfo != nil {
		this.Data["apiNodeUpgradeInfo"] = maps.Map{
			"count":   resp.ApiNodeUpgradeInfo.CountNodes,
			"version": resp.ApiNodeUpgradeInfo.NewVersion,
		}
	} else {
		this.Data["apiNodeUpgradeInfo"] = maps.Map{
			"count":   0,
			"version": "",
		}
	}
	if resp.UserNodeUpgradeInfo != nil {
		this.Data["userNodeUpgradeInfo"] = maps.Map{
			"count":   resp.UserNodeUpgradeInfo.CountNodes,
			"version": resp.UserNodeUpgradeInfo.NewVersion,
		}
	} else {
		this.Data["userNodeUpgradeInfo"] = maps.Map{
			"count":   0,
			"version": 0,
		}
	}
	if resp.AuthorityNodeUpgradeInfo != nil {
		this.Data["authorityNodeUpgradeInfo"] = maps.Map{
			"count":   resp.AuthorityNodeUpgradeInfo.CountNodes,
			"version": resp.AuthorityNodeUpgradeInfo.NewVersion,
		}
	} else {
		this.Data["authorityNodeUpgradeInfo"] = maps.Map{
			"count":   0,
			"version": "",
		}
	}
	if resp.NsNodeUpgradeInfo != nil {
		this.Data["nsNodeUpgradeInfo"] = maps.Map{
			"count":   resp.NsNodeUpgradeInfo.CountNodes,
			"version": resp.NsNodeUpgradeInfo.NewVersion,
		}
	} else {
		this.Data["nsNodeUpgradeInfo"] = maps.Map{
			"count":   0,
			"version": "",
		}
	}

	//节点数据统计 1.ddos、2.下一代防火墙、3.主机漏扫 4.web漏扫 5主机防护 6安全审计 7堡垒机
	ddos_total, _ := subassemblynode.GetNodeNum(&subassemblynode.NodeNumReq{
		Type: 1, State: "1", //
	})
	allTotal := ddos_total

	nfw_total, _ := subassemblynode.GetNodeNum(&subassemblynode.NodeNumReq{
		Type: 2, State: "1", //
	})
	allTotal += nfw_total

	//host_scan_total, _ := subassemblynode.GetNodeNum(&subassemblynode.NodeNumReq{
	//	Type: 3,State: "1", //
	//})
	//allTotal += host_scan_total

	web_scan_total, _ := subassemblynode.GetNodeNum(&subassemblynode.NodeNumReq{
		Type: 4, State: "1", //
	})
	allTotal += web_scan_total

	host_total, _ := subassemblynode.GetNodeNum(&subassemblynode.NodeNumReq{
		Type: 5, State: "1", //
	})
	allTotal += host_total

	this.Data["all_total"] = allTotal
	this.Data["node"] = maps.Map{
		"nodeCount": allTotal,
		"ddosCount": ddos_total,
		"nfwCount":  nfw_total,
		"scanCount": web_scan_total,
		"hidsCount": host_total,
	}
	//资产数据
	var scan_goal int64
	scan_goal, _ = targets.GetTargetsNum(nil)
	this.Data["assets"] = maps.Map{
		"host":      0,
		"scan_goal": scan_goal,
	}
	// 指标
	{
		var chartMaps = []maps.Map{}
		for _, chart := range resp.MetricDataCharts {
			var statMaps = []maps.Map{}
			for _, stat := range chart.MetricStats {
				statMaps = append(statMaps, maps.Map{
					"keys":  stat.Keys,
					"time":  stat.Time,
					"value": stat.Value,
					"count": stat.SumCount,
					"total": stat.SumTotal,
				})
			}
			chartMaps = append(chartMaps, maps.Map{
				"chart": maps.Map{
					"id":       chart.MetricChart.Id,
					"name":     chart.MetricChart.Name,
					"widthDiv": chart.MetricChart.WidthDiv,
					"isOn":     chart.MetricChart.IsOn,
					"maxItems": chart.MetricChart.MaxItems,
					"type":     chart.MetricChart.Type,
				},
				"item": maps.Map{
					"id":            chart.MetricChart.MetricItem.Id,
					"name":          chart.MetricChart.MetricItem.Name,
					"period":        chart.MetricChart.MetricItem.Period,
					"periodUnit":    chart.MetricChart.MetricItem.PeriodUnit,
					"valueType":     serverconfigs.FindMetricValueType(chart.MetricChart.MetricItem.Category, chart.MetricChart.MetricItem.Value),
					"valueTypeName": serverconfigs.FindMetricValueName(chart.MetricChart.MetricItem.Category, chart.MetricChart.MetricItem.Value),
					"keys":          chart.MetricChart.MetricItem.Keys,
				},
				"stats": statMaps,
			})
		}
		this.Data["metricCharts"] = chartMaps
	}

	this.Show()
}
