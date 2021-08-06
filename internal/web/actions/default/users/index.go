package users

import (
	"github.com/1uLang/zhiannet-api/common/model/edge_logins"
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
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
	Keyword string
}) {
	countResp, err := this.RPC().UserRPC().CountAllEnabledUsers(this.AdminContext(), &pb.CountAllEnabledUsersRequest{Keyword: params.Keyword})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	count := countResp.Count
	page := this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	usersResp, err := this.RPC().UserRPC().ListEnabledUsers(this.AdminContext(), &pb.ListEnabledUsersRequest{
		Keyword: params.Keyword,
		Offset:  page.Offset,
		Size:    page.Size,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//先获取所有uid
	uids, uidMap := []uint64{}, map[uint64]*edge_logins.EdgeLogins{}
	for _, v := range usersResp.Users {
		uids = append(uids, uint64(v.Id))
	}
	if len(uids) > 0 {
		//查找otp信息
		uidMap, _, err = edge_logins_server.GetListByUid(uids)
	}

	userMaps := []maps.Map{}
	for _, user := range usersResp.Users {
		var clusterMap maps.Map = nil
		if user.NodeCluster != nil {
			clusterMap = maps.Map{
				"id":   user.NodeCluster.Id,
				"name": user.NodeCluster.Name,
			}
		}

		var IsOn bool
		if len(uidMap) > 0 {
			if info, ok := uidMap[uint64(user.Id)]; ok && info.IsOn == 1 && info.Type == "otp" {
				IsOn = true
			}
		}

		userMaps = append(userMaps, maps.Map{
			"id":           user.Id,
			"username":     user.Username,
			"isOn":         user.IsOn,
			"fullname":     user.Fullname,
			"email":        user.Email,
			"mobile":       user.Mobile,
			"tel":          user.Tel,
			"createdTime":  timeutil.FormatTime("Y-m-d H:i:s", user.CreatedAt),
			"cluster":      clusterMap,
			"otpLoginIsOn": IsOn,
		})
	}
	this.Data["users"] = userMaps

	this.Show()
}
