package dns

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/dns/clusters"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/dns/domains"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeDNS)).
			//Helper(new(Helper)).
			Prefix("/dns").
			Data("teaMenu", "dns").
			Get("", new(IndexAction)).
			GetPost("/updateClusterPopup", new(UpdateClusterPopupAction)).
			Post("/providerOptions", new(ProviderOptionsAction)).
			Post("/domainOptions", new(DomainOptionsAction)).

			// 集群
			Prefix("/dns/clusters").
			Get("/cluster", new(clusters.ClusterAction)).
			Post("/sync", new(clusters.SyncAction)).

			// 域名
			Prefix("/dns/domains").
			Data("teaSubMenu", "provider").
			GetPost("/createPopup", new(domains.CreatePopupAction)).
			GetPost("/updatePopup", new(domains.UpdatePopupAction)).
			Post("/delete", new(domains.DeleteAction)).
			Post("/sync", new(domains.SyncAction)).
			Get("/routesPopup", new(domains.RoutesPopupAction)).
			GetPost("/selectPopup", new(domains.SelectPopupAction)).
			Get("/clustersPopup", new(domains.ClustersPopupAction)).
			Get("/nodesPopup", new(domains.NodesPopupAction)).
			Get("/serversPopup", new(domains.ServersPopupAction)).
			EndData().

			EndAll()
	})
}
