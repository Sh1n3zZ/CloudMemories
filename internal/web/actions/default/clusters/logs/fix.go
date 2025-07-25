// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package logs

import (
	"strings"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo/types"
)

type FixAction struct {
	actionutils.ParentAction
}

func (this *FixAction) RunPost(params struct {
	LogIds []int64
}) {
	var logIdStrings = []string{}
	for _, logId := range params.LogIds {
		logIdStrings = append(logIdStrings, types.String(logId))
	}

	defer this.CreateLogInfo(codes.NodeLog_LogFixNodeLogs, strings.Join(logIdStrings, ", "))

	_, err := this.RPC().NodeLogRPC().FixNodeLogs(this.AdminContext(), &pb.FixNodeLogsRequest{NodeLogIds: params.LogIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧数字Badge更新
	helpers.NotifyNodeLogsCountChange()

	this.Success()
}
