package nodes

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/nodes/ipAddresses"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNode)).
			Helper(new(Helper)).
			Prefix("/nodes").
			Post("/delete", new(DeleteAction)).

			// IP地址
			GetPost("/ipAddresses/createPopup", new(ipAddresses.CreatePopupAction)).
			GetPost("/ipAddresses/updatePopup", new(ipAddresses.UpdatePopupAction)).
			EndAll()
	})
}
