package transfer

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
			Helper(settingutils.NewAdvancedHelper("transfer")).
			Prefix("/settings/transfer").
			Get("", new(IndexAction)).
			Post("/validateAPI", new(ValidateAPIAction)).
			Post("/updateHosts", new(UpdateHostsAction)).
			Post("/upgradeNodes", new(UpgradeNodesAction)).
			Post("/statNodes", new(StatNodesAction)).
			EndAll()
	})
}
