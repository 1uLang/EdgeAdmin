// 主机防护使用wazuh组件
// +build wazuh

package wazuh

import (
	"fmt"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"strings"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "baseline")
}

func (this *CreateAction) RunGet(params struct{}) {

	this.Data["group"] = fmt.Sprintf("admin_%v", this.AdminId())
	this.Data["server"] = func() string {
		addr := strings.Replace(serverAddr, "http://", "", 1)
		addr = strings.Replace(addr, "https://", "", 1)
		return addr
	}
	this.Data["commands"] = maps.Map{
		"Red Hat / CentOS": maps.Map{
			"CentOS5": maps.Map{
				"i386":   "zhianhids-agent.el5.i386.rpm",
				"x86_64": "zhianhids-agent.el5.x86_64.rpm",
			},
			"CentOS6 or higher": maps.Map{
				"i386":    "zhianhids-agent.i386.rpm",
				"x86_64":  "zhianhids-agent.x86_64.rpm",
				"armhf":   "zhianhids-agent.armv7hl.rpm",
				"aarch64": "zhianhids-agent.aarch64.rpm",
			},
		},
		"Debian / Ubuntu": maps.Map{
			"i386":    "zhianhids-agent_i386.deb",
			"x86_64":  "zhianhids-agent_amd64.deb",
			"armhf":   "zhianhids-agent_armhf.deb",
			"aarch64": "zhianhids-agent_arm64.deb",
		},
		"Windows": "zhianhids-agent.msi",
		"MacOS":   "zhianhids-agent.pkg",
	}
	this.Show()
}
