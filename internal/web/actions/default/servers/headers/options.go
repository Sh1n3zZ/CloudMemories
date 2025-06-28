// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package headers

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type OptionsAction struct {
	actionutils.ParentAction
}

func (this *OptionsAction) RunPost(params struct {
	Type string
}) {
	if params.Type == "request" {
		this.Data["headers"] = serverconfigs.AllHTTPCommonRequestHeaders
	} else if params.Type == "response" {
		this.Data["headers"] = serverconfigs.AllHTTPCommonResponseHeaders
	} else {
		this.Data["headers"] = []string{}
	}

	this.Success()
}
