package users

import (
	"fmt"
	hids_user_model "github.com/1uLang/zhiannet-api/hids/model/user"
	hids_user_server "github.com/1uLang/zhiannet-api/hids/server/user"
	jumpserver_users_model "github.com/1uLang/zhiannet-api/jumpserver/model/users"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/fortcloud"
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
	moduleCodes := map[string]bool{}
	for _, code := range params.Codes {
		moduleCodes[code] = true
	}
	var user *pb.User
	//判断是否拥有了主机防护功能  有 则对应创建该用户
	if _, isExist := moduleCodes[configloaders.AdminModuleCodeHids]; isExist {

		userResp, err := this.RPC().UserRPC().FindEnabledUser(this.AdminContext(), &pb.FindEnabledUserRequest{UserId: params.UserId})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		user = userResp.User
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
	//判断是否拥有了堡垒机功能  有 则对应创建该用户
	if _, isExist := moduleCodes[configloaders.AdminModuleCodeFort]; isExist {
		if user == nil {
			userResp, err := this.RPC().UserRPC().FindEnabledUser(this.AdminContext(), &pb.FindEnabledUserRequest{UserId: params.UserId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			user = userResp.User
			if user == nil {
				this.NotFound("user", params.UserId)
				return
			}
		}
		err = fortcloud.InitAPIServer()
		if err != nil {
			this.ErrorPage(fmt.Errorf("堡垒机组件初始化失败0：%v", err))
			return
		}
		req, err := fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
		if err != nil {
			this.ErrorPage(fmt.Errorf("堡垒机组件初始化失败1：%v", err))
			return
		}
		if user.Email == "" {
			this.ErrorPage(fmt.Errorf("当前用户未绑定邮箱，请先绑定后开启堡垒机功能"))
			return
		}
		_, err = req.Users.Create(&jumpserver_users_model.CreateReq{Name: user.Username, Username: user.Username, Password: "dengbao-" + user.Username,
			Email: user.Email})
		fmt.Println("create jumpserver account : ", user.Username, "dengbao-"+user.Username)
		if err != nil {
			this.ErrorPage(fmt.Errorf("堡垒机组件同步信息失败：%v", err))
			return
		}
	}
	this.Success()
}
