package users

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/users/userutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"strings"
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
	selectList := []string{}
	//fmt.Println(allFeatures)
	//return
	idx := 0
	for _, feature := range allFeatures {
		if lists.ContainsString(userFeatureCodes, feature.Code) {
			selectList = append(selectList, feature.Code)
		}
		item := maps.Map{
			"name":        feature.Name,
			"code":        feature.Code,
			"bShowChild":  false,
			"description": feature.Description,
			"children":    []maps.Map{},
		}
		//子菜单
		if codes := strings.Split(feature.Code, "."); len(codes) == 2 {

			subItems := featureMaps[idx-1]["children"].([]maps.Map)
			subItems = append(subItems, item)
			featureMaps[idx-1]["children"] = subItems
		} else {
			featureMaps = append(featureMaps, item)
			idx++
		}
	}

	this.Data["features"] = featureMaps
	this.Data["selectList"] = selectList
	this.Data["userId"] = params.UserId
	this.Show()
}

func (this *FeaturesAction) RunPost(params struct {
	UserId int64
	Codes  []string

	Must *actions.Must
	//CSRF *actionutils.CSRF
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
	moduleCodes := map[string]bool{}
	for _, code := range params.Codes {
		moduleCodes[code] = true
	}

	this.Success()
}
