package conds

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Prefix("/servers/server/settings/conds").
			GetPost("/addGroupPopup", new(AddGroupPopupAction)).
			GetPost("/addCondPopup", new(AddCondPopupAction)).
			EndAll()
	})
}
