package server

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	adminserverutils "github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/settings/server/admin-server-utils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	this.Data["serverIsChanged"] = adminserverutils.ServerConfigIsChanged

	serverConfig, err := adminserverutils.LoadServerConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["serverConfig"] = serverConfig

	this.Show()
}
