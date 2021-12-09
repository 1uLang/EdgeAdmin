// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build plus
// +build plus

package boardutils

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

func InitBoard(parent *actionutils.ParentAction) error {
	countResp, err := parent.RPC().NodeLogRPC().CountAllUnreadNodeLogs(parent.AdminContext(), &pb.CountAllUnreadNodeLogsRequest{})
	if err != nil {
		return err
	}
	parent.Data["countEvents"] = countResp.Count
	return nil
}
