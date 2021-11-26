package hostlist

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Prefix("/hostlist").
			GetPost("", new(IndexAction)).
			Prefix("/hostlist/create").
			GetPost("", new(CreatePopupAction)).
			Prefix("/hostlist/activity").
			GetPost("", new(ActivityAction)).
			Prefix("/hostlist/spec").
			GetPost("", new(SpecAction)).
			Prefix("/hostlist/image").
			GetPost("", new(ImageAction)).
			Prefix("/hostlist/network").
			GetPost("", new(NetworkAction)).
			Prefix("/hostlist/disk").
			GetPost("", new(DiskAction)).
			Prefix("/hostlist/candidate").
			GetPost("", new(CandidateAction)).
			EndAll()
	})
}
