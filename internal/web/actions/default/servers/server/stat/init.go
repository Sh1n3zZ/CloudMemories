package stat

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/serverutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(serverutils.NewServerHelper()).
			Prefix("/servers/server/stat").
			Get("", new(IndexAction)).
			Get("/hourlyRequests", new(HourlyRequestsAction)).
			Get("/dailyRequests", new(DailyRequestsAction)).
			Get("/regions", new(RegionsAction)).
			Get("/providers", new(ProvidersAction)).
			Get("/clients", new(ClientsAction)).
			Get("/waf", new(WafAction)).
			EndAll()
	})
}
