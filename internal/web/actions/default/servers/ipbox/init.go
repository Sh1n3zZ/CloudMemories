// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ipbox

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Prefix("/servers/ipbox").
			Get("", new(IndexAction)).
			Post("/addIP", new(AddIPAction)).
			Post("/deleteFromList", new(DeleteFromListAction)).
			EndAll()
	})
}
