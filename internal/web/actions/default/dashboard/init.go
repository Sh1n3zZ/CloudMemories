package dashboard

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.Prefix("/dashboard").
			Data("teaMenu", "dashboard").
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeCommon)).
			GetPost("", new(IndexAction)).
			Post("/restartLocalAPINode", new(RestartLocalAPINodeAction)).
			EndAll()
	})
}
