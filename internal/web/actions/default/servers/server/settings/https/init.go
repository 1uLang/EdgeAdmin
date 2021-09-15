package https

import (
	"encoding/json"
	"github.com/1uLang/zhiannet-api/common/server/attack_check_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/serverutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/iwind/TeaGo"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/types"
)

var chanWAFServer chan uint64

func init() {

	//设置 waf chan
	chanWAFServer = make(chan uint64, 1)
	attack_check_server.SetChan(chanWAFServer)

	go selectWAFServerChan()

	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(serverutils.NewServerHelper()).
			Data("teaMenu", "waf").
			Prefix("/servers/server/settings/https").
			GetPost("", new(IndexAction)).
			GetPost("/requestCertPopup", new(RequestCertPopupAction)).
			EndAll()
	})
}
func selectWAFServerChan() {

	for {
		select {
		case serverId := <-chanWAFServer:
			go getServerInfo(int64(serverId))
		}
	}

}
func getServerInfo(serverId int64) {
	this := new(IndexAction)

	//get
	server, _, isOk := serverutils.FindServer(this.Parent(), serverId)
	if !isOk {
		return
	}
	httpsConfig := &serverconfigs.HTTPSProtocolConfig{}
	if len(server.HttpsJSON) > 0 {
		err := json.Unmarshal(server.HttpsJSON, httpsConfig)
		if err != nil {
			logs.Errorf(string(server.HttpsJSON) + " json unmarshal error : " + err.Error())
			return
		}
	} else {
		httpsConfig.IsOn = true
	}

	var sslPolicy *sslconfigs.SSLPolicy
	if httpsConfig.SSLPolicyRef == nil || httpsConfig.SSLPolicyRef.SSLPolicyId <= 0 {
		return
	}
	sslPolicyConfigResp, err := this.RPC().SSLPolicyRPC().FindEnabledSSLPolicyConfig(this.AdminContext(), &pb.FindEnabledSSLPolicyConfigRequest{SslPolicyId: httpsConfig.SSLPolicyRef.SSLPolicyId})
	if err != nil {
		logs.Errorf("FindEnabledSSLPolicyConfig error : " + err.Error())
		return
	}
	sslPolicyConfigJSON := sslPolicyConfigResp.SslPolicyJSON
	if len(sslPolicyConfigJSON) <= 0 {
		return
	}
	sslPolicy = &sslconfigs.SSLPolicy{}
	err = json.Unmarshal(sslPolicyConfigJSON, sslPolicy)
	if err != nil {
		logs.Errorf(string(sslPolicyConfigJSON) + " json unmarshal error : " + err.Error())
		return
	}

	certsJSON, err := json.Marshal(sslPolicy.CertRefs)
	if err != nil {
		logs.Errorf("sslPolicy CertRefs json marshal error : " + err.Error())
		return
	}

	hstsJSON, err := json.Marshal(sslPolicy.HSTS)
	if err != nil {
		logs.Errorf("sslPolicy HSTS json marshal error : " + err.Error())
		return
	}

	clientCACertsJSON, err := json.Marshal(sslPolicy.ClientCARefs)
	if err != nil {
		logs.Errorf("sslPolicy ClientCARefs json marshal error : " + err.Error())
		return
	}
	//update
	_, err = this.RPC().SSLPolicyRPC().UpdateSSLPolicy(this.AdminContext(), &pb.UpdateSSLPolicyRequest{
		SslPolicyId:       sslPolicy.Id,
		Http2Enabled:      sslPolicy.HTTP2Enabled,
		MinVersion:        sslPolicy.MinVersion,
		SslCertsJSON:      certsJSON,
		HstsJSON:          hstsJSON,
		ClientAuthType:    types.Int32(sslPolicy.ClientAuthType),
		ClientCACertsJSON: clientCACertsJSON,
		CipherSuitesIsOn:  sslPolicy.CipherSuitesIsOn,
		CipherSuites:      sslPolicy.CipherSuites,
	})
	if err != nil {
		logs.Errorf("UpdateSSLPolicy error : " + err.Error())
		return
	}
}
