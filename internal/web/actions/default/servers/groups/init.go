package groups

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/groups/group"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(NewHelper()).
			Data("teaMenu", "servers").
			Data("teaSubMenu", "group").
			Prefix("/servers/groups").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/selectPopup", new(SelectPopupAction)).
			Post("/sort", new(SortAction)).

			// 详情
			Prefix("/servers/groups/group").
			Get("", new(group.IndexAction)).
			Post("/delete", new(group.DeleteAction)).
			GetPost("/update", new(group.UpdateAction)).
			EndAll()
	})
}
