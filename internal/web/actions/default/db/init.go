package db

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/settings/settingutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeSetting)).
			Helper(new(Helper)).
			Helper(settingutils.NewAdvancedHelper("dbNodes")).
			Prefix("/db").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/update", new(UpdateAction)).
			Post("/delete", new(DeleteAction)).
			GetPost("/clean", new(CleanAction)).
			Post("/deleteTable", new(DeleteTableAction)).
			Post("/truncateTable", new(TruncateTableAction)).
			Get("/node", new(NodeAction)).
			Get("/logs", new(LogsAction)).
			Post("/status", new(StatusAction)).
			EndAll()
	})
}
