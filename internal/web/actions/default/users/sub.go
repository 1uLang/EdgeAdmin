package users

import (
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
	"github.com/1uLang/zhiannet-api/common/server/channels_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_users_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type SubAction struct {
	actionutils.ParentAction
}

func (this *SubAction) Init() {
	this.Nav("", "", "")
}

func (this *SubAction) RunGet(params struct {
	Keyword  string
	PageNum  int
	PageSize int
	ParentId uint64
}) {
	//当前用户信息
	user, err := edge_users_server.GetUserInfo(uint64(params.ParentId))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	countAccessKeyResp, err := this.RPC().UserAccessKeyRPC().CountAllEnabledUserAccessKeys(this.AdminContext(), &pb.CountAllEnabledUserAccessKeysRequest{UserId: int64(params.ParentId)})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["user"] = maps.Map{
		"id":              user.ID,
		"username":        user.Username,
		"fullname":        user.Fullname,
		"email":           user.Email,
		"tel":             user.Tel,
		"remark":          user.Remark,
		"mobile":          user.Mobile,
		"countAccessKeys": countAccessKeyResp.Count,
	}

	if params.PageSize == 0 {
		params.PageSize = 20
	}
	if params.PageNum == 0 {
		params.PageNum = 1
	}
	list, total, err := edge_users_server.GetList(&edge_users.ListReq{
		PageNum:  params.PageNum,
		PageSize: params.PageSize,
		Username: params.Keyword,
		ParentId: params.ParentId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	page := this.NewPage(total)
	this.Data["page"] = page.AsHTML()

	userMaps := []maps.Map{}
	if len(list) > 0 {
		for _, user := range list {
			userMaps = append(userMaps, maps.Map{
				"id":           user.ID,
				"username":     user.Username,
				"isOn":         user.Ison,
				"fullname":     user.Fullname,
				"email":        user.Email,
				"mobile":       user.Mobile,
				"tel":          user.Tel,
				"createdTime":  timeutil.FormatTime("Y-m-d H:i:s", user.Createdat),
				"cluster":      user.Cluster,
				"channel":      user.Channels,
				"otpLoginIsOn": user.OtpLoginIsOn,
				"subTotal":     user.SubTotal,
			})
		}

	}
	this.Data["users"] = userMaps

	//渠道列表
	channels, _, err := channels_server.GetList(&channels.ChannelReq{
		//Status:   "1",
		PageNum:  1,
		PageSize: 999,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	chanList := make([]maps.Map, 0)
	chanList = append(chanList, maps.Map{
		"id": 0, "name": "全部",
	})
	if len(channels) > 0 {
		for _, v := range channels {
			chanList = append(chanList, maps.Map{
				"id": v.Id, "name": v.Name,
			})
		}
	}
	this.Data["channels"] = chanList
	this.Show()
}
