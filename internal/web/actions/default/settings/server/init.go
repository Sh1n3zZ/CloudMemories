package server

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/settings/settingutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(settingutils.NewHelper("server")).
			Prefix("/settings/server").
			Get("", new(IndexAction)).
			GetPost("/updateHTTPPopup", new(UpdateHTTPPopupAction)).
			GetPost("/updateHTTPSPopup", new(UpdateHTTPSPopupAction)).
			EndAll()
	})
}
