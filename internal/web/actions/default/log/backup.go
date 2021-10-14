package log

import (
	"bytes"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/lists"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"github.com/tealeg/xlsx/v3"
	"strings"
)

type BackupAction struct {
	actionutils.ParentAction
}

func (this *BackupAction) Init() {
	this.Nav("", "", "")
}

func (this *BackupAction) RunGet(params struct {
	DayFrom  string
	DayTo    string
	Keyword  string
	UserType string
}) {
	logsResp, err := this.RPC().LogRPC().ListLogs(this.AdminContext(), &pb.ListLogsRequest{
		Offset:   0,
		Size:     1000, // 日志最大导出1000条，TODO 将来可以配置
		DayFrom:  params.DayFrom,
		DayTo:    params.DayTo,
		Keyword:  params.Keyword,
		UserType: params.UserType,
	})
	if err != nil {
		this.FailField("error", err.Error())
		return
	}

	wb := xlsx.NewFile()
	sheet, err := wb.AddSheet("default")
	if err != nil {
		this.FailField("error", err.Error())
		return
	}

	// 头部
	{
		row := sheet.AddRow()
		row.SetHeight(25)
		row.AddCell().SetString("ID")
		row.AddCell().SetString("日期")
		row.AddCell().SetString("用户")
		row.AddCell().SetString("描述")
		row.AddCell().SetString("IP")
		row.AddCell().SetString("区域")
		row.AddCell().SetString("运营商")
		row.AddCell().SetString("页面地址")
	}

	// 数据
	for _, log := range logsResp.Logs {
		regionName := ""
		ispName := ""
		regionResp, err := this.RPC().IPLibraryRPC().LookupIPRegion(this.AdminContext(), &pb.LookupIPRegionRequest{Ip: log.Ip})
		if err != nil {
			this.FailField("error", err.Error())
			return
		}
		if regionResp.IpRegion != nil {
			pieces := []string{}
			if len(regionResp.IpRegion.Country) > 0 {
				pieces = append(pieces, regionResp.IpRegion.Country)
			}
			if len(regionResp.IpRegion.Province) > 0 && !lists.ContainsString(pieces, regionResp.IpRegion.Province) {
				pieces = append(pieces, regionResp.IpRegion.Province)
			}
			if len(regionResp.IpRegion.City) > 0 && !lists.ContainsString(pieces, regionResp.IpRegion.City) && !lists.ContainsString(pieces, strings.TrimSuffix(regionResp.IpRegion.Province, "市")) {
				pieces = append(pieces, regionResp.IpRegion.City)
			}
			regionName = strings.Join(pieces, " ")

			if len(regionResp.IpRegion.Isp) > 0 {
				ispName = regionResp.IpRegion.Isp
			}
		}

		row := sheet.AddRow()
		row.SetHeight(25)
		row.AddCell().SetInt64(log.Id)
		row.AddCell().SetString(timeutil.FormatTime("Y-m-d H:i:s", log.CreatedAt))
		if log.UserId > 0 {
			row.AddCell().SetString("用户 | " + log.UserName)
		} else {
			row.AddCell().SetString(log.UserName)
		}
		row.AddCell().SetString(log.Description)
		row.AddCell().SetString(log.Ip)
		row.AddCell().SetString(regionName)
		row.AddCell().SetString(ispName)
		row.AddCell().SetString(log.Action)
	}
	timeStr := timeutil.Format("YmdHis")
	//this.AddHeader("Content-Type", "application/vnd.ms-excel")
	//this.AddHeader("Content-Disposition", "attachment; filename=\"LOG-"+timeStr+".xlsx\"")
	//this.AddHeader("Cache-Control", "max-age=0")

	buf := bytes.NewBuffer([]byte{})
	err = wb.Write(buf)
	if err != nil {
		this.FailField("error", err.Error())
		return
	}

	//this.AddHeader("Content-Length", strconv.Itoa(buf.Len()))
	//this.Write(buf.Bytes())

	// 获取token
	token, err := model.QueryTokenByUID(this.AdminId(), 1)
	if err != nil {
		this.FailField("error", err.Error())
		return
	}

	name := "LOG-" + timeStr + ".xlsx"
	// err = request.UploadFile(token, name, bytes.NewReader(upFile))
	request.CreateFoler(token, "", "日志备份")
	err = request.UploadFileWithPath(token, name, bytes.NewReader(buf.Bytes()), "/remote.php/dav/files/admin/日志备份/")
	if err != nil {
		this.FailField("error", err.Error())
		return
	}
	this.Success()
}
