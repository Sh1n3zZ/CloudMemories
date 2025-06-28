package tcpReverseProxy

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Data("teaMenu", "servers").
			Data("teaSubMenu", "group").
			Prefix("/servers/groups/group/settings/tcpReverseProxy").
			Get("", new(IndexAction)).
			GetPost("/scheduling", new(SchedulingAction)).
			GetPost("/setting", new(SettingAction)).
			EndAll()
	})
}
