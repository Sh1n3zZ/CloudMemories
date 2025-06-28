// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package setup

import "github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"

var currentStatusText = ""

type StatusAction struct {
	actionutils.ParentAction
}

func (this *StatusAction) RunPost(params struct{}) {
	this.Data["statusText"] = currentStatusText
	this.Success()
}
