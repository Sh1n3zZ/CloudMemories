// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package iplists

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type LevelOptionsAction struct {
	actionutils.ParentAction
}

func (this *LevelOptionsAction) RunPost(params struct{}) {
	this.Data["levels"] = firewallconfigs.FindAllFirewallEventLevels()

	this.Success()
}
