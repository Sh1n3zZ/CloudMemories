// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ddosProtection

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/messageconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/nodes/nodeutils"
)

type StatusAction struct {
	actionutils.ParentAction
}

func (this *StatusAction) RunPost(params struct {
	NodeId int64
}) {
	results, err := nodeutils.SendMessageToNodeIds(this.AdminContext(), []int64{params.NodeId}, messageconfigs.MessageCodeCheckLocalFirewall, &messageconfigs.CheckLocalFirewallMessage{
		Name: "nftables",
	}, 10)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["results"] = results
	this.Success()
}
