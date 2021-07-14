package users

import (
	"fmt"
	hids_user_model "github.com/1uLang/zhiannet-api/hids/model/user"
	hids_user_server "github.com/1uLang/zhiannet-api/hids/server/user"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/hids"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/users/userutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
)

type FeaturesAction struct {
	actionutils.ParentAction
}

func (this *FeaturesAction) Init() {
	this.Nav("", "", "feature")
}

func (this *FeaturesAction) RunGet(params struct {
	UserId int64
}) {
	err := userutils.InitUser(this.Parent(), params.UserId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	featuresResp, err := this.RPC().UserRPC().FindAllUserFeatureDefinitions(this.AdminContext(), &pb.FindAllUserFeatureDefinitionsRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	allFeatures := featuresResp.Features

	userFeaturesResp, err := this.RPC().UserRPC().FindUserFeatures(this.AdminContext(), &pb.FindUserFeaturesRequest{UserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	userFeatureCodes := []string{}
	for _, userFeature := range userFeaturesResp.Features {
		userFeatureCodes = append(userFeatureCodes, userFeature.Code)
	}

	featureMaps := []maps.Map{}
	for _, feature := range allFeatures {
		featureMaps = append(featureMaps, maps.Map{
			"name":        feature.Name,
			"code":        feature.Code,
			"description": feature.Description,
			"isChecked":   lists.ContainsString(userFeatureCodes, feature.Code),
		})
	}

	this.Data["features"] = featureMaps

	this.Show()
}

func (this *FeaturesAction) RunPost(params struct {
	UserId int64
	Codes  []string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("设置用户 %d 的功能列表", params.UserId)

	_, err := this.RPC().UserRPC().UpdateUserFeatures(this.AdminContext(), &pb.UpdateUserFeaturesRequest{
		UserId:       params.UserId,
		FeatureCodes: params.Codes,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//判断是否拥有了主机防护功能  有 则对应创建该用户
	var hasHids bool
	for _, code := range params.Codes {
		if code == configloaders.AdminModuleCodeHids {
			hasHids = true
			break
		}
	}
	if hasHids {

		userResp, err := this.RPC().UserRPC().FindEnabledUser(this.AdminContext(), &pb.FindEnabledUserRequest{UserId: params.UserId})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		user := userResp.User
		if user == nil {
			this.NotFound("user", params.UserId)
			return
		}
		err = hids.InitAPIServer()
		if err != nil {
			this.ErrorPage(fmt.Errorf("主机防护组件初始化失败：%v", err))
			return
		}
		_, err = hids_user_server.Add(&hids_user_model.AddReq{UserName: user.Username, Password: "dengbao-" + user.Username, Role: 3})
		if err != nil {
			this.ErrorPage(fmt.Errorf("主机防护组件同步信息失败：%v", err))
			return
		}
	}
	this.Success()
}
