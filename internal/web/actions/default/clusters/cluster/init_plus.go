//go:build plus
// +build plus

package cluster

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/cluster/boards"
	nodeboards "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/cluster/node/boards"
	clusters "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/clusterutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNode)).
			Helper(clusters.NewClusterHelper()).
			Data("teaMenu", "clusters").

			// 节点相关
			Prefix("/clusters/cluster/node").
			GetPost("/boards", new(nodeboards.IndexAction)).
			Post("/boards/data", new(nodeboards.DataAction)).

			// 看板相关
			Prefix("/clusters/cluster/boards").
			GetPost("", new(boards.IndexAction)).
			EndAll()
	})
}
