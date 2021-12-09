// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package traffic

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("traffic")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	this.Data["serverId"] = params.ServerId

	configResp, err := this.RPC().ServerRPC().FindEnabledServerTrafficLimit(this.AdminContext(), &pb.FindEnabledServerTrafficLimitRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	var trafficLimitConfig = &serverconfigs.TrafficLimitConfig{}
	if len(configResp.TrafficLimitJSON) > 0 {
		err = json.Unmarshal(configResp.TrafficLimitJSON, trafficLimitConfig)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	this.Data["trafficLimitConfig"] = trafficLimitConfig

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId         int64
	TrafficLimitJSON []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("修改服务 %d 流量限制", params.ServerId)

	_, err := this.RPC().ServerRPC().UpdateServerTrafficLimit(this.AdminContext(), &pb.UpdateServerTrafficLimitRequest{
		ServerId:         params.ServerId,
		TrafficLimitJSON: params.TrafficLimitJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
