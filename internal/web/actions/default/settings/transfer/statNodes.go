// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package transfer

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type StatNodesAction struct {
	actionutils.ParentAction
}

func (this *StatNodesAction) RunPost(params struct{}) {
	countNodesResp, err := this.RPC().NodeRPC().CountAllEnabledNodesMatch(this.AdminContext(), &pb.CountAllEnabledNodesMatchRequest{ActiveState: 1})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["countNodes"] = countNodesResp.Count

	this.Success()
}
