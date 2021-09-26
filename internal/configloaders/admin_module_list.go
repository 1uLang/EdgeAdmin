package configloaders

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/systemconfigs"
	"strings"
)

type AdminModuleList struct {
	IsSuper  bool
	Modules  []*systemconfigs.AdminModule
	Fullname string
	Theme    string
}

func (this *AdminModuleList) Allow(module string, memu ...bool) bool {
	if this.IsSuper {
		return true
	}
	for _, m := range this.Modules {
		if m.Code == module || (strings.HasPrefix(m.Code, module) && (len(memu) == 0 || !memu[0])) {
			return true
		}
	}
	return false
}
