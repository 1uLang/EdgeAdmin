package resmon

import (
	"regexp"

	"github.com/1uLang/zhiannet-api/resmon/request"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {

	this.Data["list"] = []int{}
	this.Data["total"] = 0
	defer this.Show()
	al, err := request.AgentList()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}

	this.Data["list"] = al.List
	this.Data["total"] = al.Total
}

func (this *IndexAction) RunPost(params struct {
	Name string
	Host string
	Aid  uint8
	On   bool
	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入主机名").
		Field("host", params.Host).
		Require("请输入主机地址").
		Field("aid", params.Aid).
		Require("请输入主机类型")

	reg := regexp.MustCompile(`^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$`)
	if !reg.MatchString(params.Host) {
		this.Fail("请输入正确的ip地址")
	}

	err := request.AddAgent(params.Name, params.Host, params.On, params.Aid)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
