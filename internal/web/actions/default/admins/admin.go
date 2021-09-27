package admins

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

type AdminAction struct {
	actionutils.ParentAction
}

func (this *AdminAction) Init() {
	this.Nav("", "", "index")
}

func (this *AdminAction) RunGet(params struct {
	AdminId int64
}) {
	adminResp, err := this.RPC().AdminRPC().FindEnabledAdmin(this.AdminContext(), &pb.FindEnabledAdminRequest{AdminId: params.AdminId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	admin := adminResp.Admin
	if admin == nil {
		this.NotFound("admin", params.AdminId)
		return
	}

	// AccessKey数量
	countAccessKeyResp, err := this.RPC().UserAccessKeyRPC().CountAllEnabledUserAccessKeys(this.AdminContext(), &pb.CountAllEnabledUserAccessKeysRequest{AdminId: params.AdminId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	countAccessKeys := countAccessKeyResp.Count

	role := 0
	if !admin.IsSuper {
		if admin.Modules[0].Code == roleMenus[1][0] {
			role = 1
		} else if admin.Modules[0].Code == roleMenus[2][0] {
			role = 2
		} else {
			role = 3
		}
	}
	this.Data["admin"] = maps.Map{
		"id":              admin.Id,
		"fullname":        admin.Fullname,
		"username":        admin.Username,
		"isOn":            admin.IsOn,
		"isSuper":         admin.IsSuper,
		"canLogin":        admin.CanLogin,
		"countAccessKeys": countAccessKeys,
		"role":            role,
	}

	// 权限
	moduleMaps := []maps.Map{}
	for _, m := range configloaders.AllModuleMaps() {
		code := m.GetString("code")
		isChecked := false
		for _, module := range admin.Modules {
			if module.Code == code {
				isChecked = true
				break
			}
		}
		if isChecked {
			moduleMaps = append(moduleMaps, m)
		}
	}
	this.Data["modules"] = moduleMaps

	// OTP
	this.Data["otp"] = nil
	if admin.OtpLogin != nil && admin.OtpLogin.IsOn {
		loginParams := maps.Map{}
		err = json.Unmarshal(admin.OtpLogin.ParamsJSON, &loginParams)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		this.Data["otp"] = maps.Map{
			"isOn":   true,
			"params": loginParams,
		}
	}

	this.Show()
}
