package users

import (
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
	"github.com/1uLang/zhiannet-api/common/server/channels_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_users_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Keyword    string
	SelectChan uint64
}) {

	page := this.NewPage(0)

	list, total, err := edge_users_server.GetList(&edge_users.ListReq{
		PageNum:   int(page.Current),
		PageSize:  int(page.Size),
		Username:  params.Keyword,
		ChannelId: params.SelectChan,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	page = this.NewPage(total)
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
	this.Data["selectChan"] = params.SelectChan
	this.Show()
}

//
//func (this *IndexAction) RunGet(params struct {
//	Keyword string
//}) {
//	countResp, err := this.RPC().UserRPC().CountAllEnabledUsers(this.AdminContext(), &pb.CountAllEnabledUsersRequest{Keyword: params.Keyword})
//	if err != nil {
//		this.ErrorPage(err)
//		return
//	}
//	count := countResp.Count
//	page := this.NewPage(count)
//	this.Data["page"] = page.AsHTML()
//
//	usersResp, err := this.RPC().UserRPC().ListEnabledUsers(this.AdminContext(), &pb.ListEnabledUsersRequest{
//		Keyword: params.Keyword,
//		Offset:  page.Offset,
//		Size:    page.Size,
//	})
//	if err != nil {
//		this.ErrorPage(err)
//		return
//	}
//	//先获取所有uid
//	uids, uidMap := []uint64{}, map[uint64]*edge_logins.EdgeLogins{}
//	for _, v := range usersResp.Users {
//		uids = append(uids, uint64(v.Id))
//	}
//	if len(uids) > 0 {
//		//查找otp信息
//		uidMap, _, err = edge_logins_server.GetListByUid(uids)
//	}
//
//	userMaps := []maps.Map{}
//	for _, user := range usersResp.Users {
//		var clusterMap maps.Map = nil
//		if user.NodeCluster != nil {
//			clusterMap = maps.Map{
//				"id":   user.NodeCluster.Id,
//				"name": user.NodeCluster.Name,
//			}
//		}
//
//		var IsOn bool
//		if len(uidMap) > 0 {
//			if info, ok := uidMap[uint64(user.Id)]; ok && info.IsOn == 1 && info.Type == "otp" {
//				IsOn = true
//			}
//		}
//
//		userMaps = append(userMaps, maps.Map{
//			"id":           user.Id,
//			"username":     user.Username,
//			"isOn":         user.IsOn,
//			"fullname":     user.Fullname,
//			"email":        user.Email,
//			"mobile":       user.Mobile,
//			"tel":          user.Tel,
//			"createdTime":  timeutil.FormatTime("Y-m-d H:i:s", user.CreatedAt),
//			"cluster":      clusterMap,
//			"otpLoginIsOn": IsOn,
//		})
//	}
//	this.Data["users"] = userMaps
//
//	this.Show()
//}
