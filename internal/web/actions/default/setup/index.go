package setup

import (
	"net"
	"strings"

	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	var currentHost = this.Request.Host
	if strings.Contains(this.Request.Host, ":") {
		host, _, err := net.SplitHostPort(this.Request.Host)
		if err == nil {
			currentHost = host
		}
	}
	if net.ParseIP(currentHost) != nil && currentHost != "localhost" && currentHost != "127.0.0.1" {
		this.Data["currentHost"] = currentHost
	} else {
		this.Data["currentHost"] = ""
	}

	this.Show()
}
