package origins

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Prefix("/servers/server/settings/origins").
			GetPost("/addPopup", new(AddPopupAction)).
			Post("/delete", new(DeleteAction)).
			GetPost("/updatePopup", new(UpdatePopupAction)).
			EndAll()
	})
}
