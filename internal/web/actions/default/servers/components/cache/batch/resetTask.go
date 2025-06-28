// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package cache

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type ResetTaskAction struct {
	actionutils.ParentAction
}

func (this *ResetTaskAction) RunPost(params struct {
	TaskId int64
}) {
	this.CreateLogInfo(codes.HTTPCacheTask_LogResetHTTPCacheTask, params.TaskId)

	_, err := this.RPC().HTTPCacheTaskRPC().ResetHTTPCacheTask(this.AdminContext(), &pb.ResetHTTPCacheTaskRequest{HttpCacheTaskId: params.TaskId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
