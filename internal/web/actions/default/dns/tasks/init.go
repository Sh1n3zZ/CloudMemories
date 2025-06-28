package tasks

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/clusters/clusterutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeDNS)).
			Helper(clusterutils.NewClustersHelper()).
			Prefix("/dns/tasks").
			GetPost("/listPopup", new(ListPopupAction)).
			Post("/check", new(CheckAction)).
			Post("/delete", new(DeleteAction)).
			Post("/deleteAll", new(DeleteAllAction)).
			EndAll()
	})
}
