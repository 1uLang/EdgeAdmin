// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build plus
// +build plus

package dashboard

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/dashboard/boards"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Data("teaMenu", "dashboard").
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeCommon)).

			// 看板
			Prefix("/dashboard/boards").
			GetPost("", new(boards.IndexAction)).
			Get("/waf", new(boards.WafAction)).
			Post("/wafLogs", new(boards.WafLogsAction)).
			Get("/dns", new(boards.DnsAction)).
			Get("/user", new(boards.UserAction)).
			Get("/events", new(boards.EventsAction)).
			Post("/readLogs", new(boards.ReadLogsAction)).
			Post("/readAllLogs", new(boards.ReadAllLogsAction)).
			EndAll()
	})
}
