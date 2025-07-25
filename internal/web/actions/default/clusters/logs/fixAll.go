// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package logs

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
)

type FixAllAction struct {
	actionutils.ParentAction
}

func (this *FixAllAction) RunPost(params struct {
}) {
	defer this.CreateLogInfo(codes.NodeLog_LogFixAllLogs)

	_, err := this.RPC().NodeLogRPC().FixAllNodeLogs(this.AdminContext(), &pb.FixAllNodeLogsRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧数字Badge更新
	helpers.NotifyNodeLogsCountChange()

	this.Success()
}
