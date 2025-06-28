package updates

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
			Helper(settingutils.NewHelper("updates")).
			Prefix("/settings/updates").
			GetPost("", new(IndexAction)).
			Post("/update", new(UpdateAction)).
			Post("/ignoreVersion", new(IgnoreVersionAction)).
			Post("/resetIgnoredVersion", new(ResetIgnoredVersionAction)).
			GetPost("/upgrade", new(UpgradeAction)).
			EndAll()
	})
}
