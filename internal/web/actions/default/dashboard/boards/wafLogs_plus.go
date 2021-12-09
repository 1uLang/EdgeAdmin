// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build plus
// +build plus

package boards

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type WafLogsAction struct {
	actionutils.ParentAction
}

func (this *WafLogsAction) RunPost(params struct{}) {
	resp, err := this.RPC().HTTPAccessLogRPC().ListHTTPAccessLogs(this.AdminContext(), &pb.ListHTTPAccessLogsRequest{
		HasFirewallPolicy: true,
		Reverse:           false,
		Day:               timeutil.Format("Ymd"),
		Size:              5,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["accessLogs"] = resp.HttpAccessLogs
	this.Success()
}
