package channels

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"github.com/1uLang/zhiannet-api/common/server/channels_server"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
	"io"
	"strings"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "")
}

func (this *CreateAction) RunGet(params struct {
	Id uint64
}) {

	info, _ := channels_server.GetInfo(params.Id)
	this.Data["name"] = info.Name
	this.Data["user"] = info.User
	this.Data["mobile"] = info.Mobile
	this.Data["productName"] = info.ProductName
	this.Data["status"] = info.Status
	this.Data["domain"] = info.Domain
	this.Data["remake"] = info.Remake
	this.Data["logo"] = fmt.Sprintf("/ui/image/%v", info.Logo)
	this.Data["id"] = info.Id
	if params.Id == 0 {
		this.Data["status"] = 1
	}
	this.Show()
}

func (this *CreateAction) RunPost(params struct {
	Id          uint64
	Name        string
	Domain      string
	ProductName string
	Logo        *actions.File
	Mobile      string
	User        string
	Remake      string
	Status      int
	OldLogo     string
	Must        *actions.Must
	CSRF        *actionutils.CSRF
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入渠道名称")
	// 上传Logo文件
	logoname := ""
	if params.Logo != nil {
		createResp, err := this.RPC().FileRPC().CreateFile(this.AdminContext(), &pb.CreateFileRequest{
			Filename: params.Logo.Filename,
			Size:     params.Logo.Size,
			IsPublic: true,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		fileId := createResp.FileId

		// 上传内容
		buf := make([]byte, 512*1024)
		reader, err := params.Logo.OriginFile.Open()
		if err != nil {
			this.ErrorPage(err)
			return
		}
		for {
			n, err := reader.Read(buf)
			if n > 0 {
				_, err = this.RPC().FileChunkRPC().CreateFileChunk(this.AdminContext(), &pb.CreateFileChunkRequest{
					FileId: fileId,
					Data:   buf[:n],
				})
				if err != nil {
					this.Fail("上传失败：" + err.Error())
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				this.Fail("上传失败：" + err.Error())
			}
		}

		// 置为已完成
		_, err = this.RPC().FileRPC().UpdateFileFinished(this.AdminContext(), &pb.UpdateFileFinishedRequest{FileId: fileId})
		if err != nil {
			this.ErrorPage(err)
		}
		logoname = fmt.Sprintf("%v", fileId)
	} else {
		if params.OldLogo == "" {
			logoname = ""
		} else {
			logoname = strings.Replace(params.OldLogo, "/ui/image/", "", 1)
		}
	}

	req := &channels.Channels{
		Name:        params.Name,
		User:        params.User,
		Mobile:      params.Mobile,
		ProductName: params.ProductName,
		Status:      params.Status,
		Domain:      params.Domain,
		Logo:        logoname,
		Remake:      params.Remake,
		Id:          params.Id,
	}

	id, err := channels_server.Add(req)
	if err != nil {
		this.Fail(err.Error())
		return
	}
	defer this.CreateLogInfo("创建渠道 %d", id)

	this.Success()
}
