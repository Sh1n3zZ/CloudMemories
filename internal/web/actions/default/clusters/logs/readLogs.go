// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package logs

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/nodeconfigs"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
)

type ReadLogsAction struct {
	actionutils.ParentAction
}

func (this *ReadLogsAction) RunPost(params struct {
	LogIds []int64

	NodeId int64
}) {
	_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{
		NodeLogIds: params.LogIds,
		NodeId:     params.NodeId,
		Role:       nodeconfigs.NodeRoleNode,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧数字Badge更新
	helpers.NotifyNodeLogsCountChange()

	this.Success()
}
