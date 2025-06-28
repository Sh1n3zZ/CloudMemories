// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dns

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/rands"
)

type RandomNameAction struct {
	actionutils.ParentAction
}

func (this *RandomNameAction) RunPost(params struct{}) {
	this.Data["name"] = "cluster" + rands.HexString(8)

	this.Success()
}
