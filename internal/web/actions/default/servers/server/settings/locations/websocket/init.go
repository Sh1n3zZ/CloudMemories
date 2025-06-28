package websocket

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/server/settings/locations/locationutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/serverutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(locationutils.NewLocationHelper()).
			Helper(serverutils.NewServerHelper()).
			Data("tinyMenuItem", "websocket").
			Prefix("/servers/server/settings/locations/websocket").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
