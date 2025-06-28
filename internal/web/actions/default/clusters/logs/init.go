package logs

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNode)).
			Data("teaMenu", "clusters").
			Data("teaSubMenu", "log").
			Prefix("/clusters/logs").
			Get("", new(IndexAction)).
			Post("/readLogs", new(ReadLogsAction)).
			Post("/readAllLogs", new(ReadAllLogsAction)).
			Post("/fix", new(FixAction)).
			Post("/fixAll", new(FixAllAction)).
			Post("/deleteAll", new(DeleteAllAction)).
			EndAll()
	})
}
