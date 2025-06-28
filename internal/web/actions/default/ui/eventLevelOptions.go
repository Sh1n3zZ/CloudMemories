package ui

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type EventLevelOptionsAction struct {
	actionutils.ParentAction
}

func (this *EventLevelOptionsAction) RunPost(params struct{}) {
	this.Data["eventLevels"] = firewallconfigs.FindAllFirewallEventLevels()

	this.Success()
}
