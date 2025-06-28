// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package cache

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/rpc"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

func InitMenu(parent *actionutils.ParentAction) error {
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		return err
	}

	countTasksResp, err := rpcClient.HTTPCacheTaskRPC().CountDoingHTTPCacheTasks(parent.AdminContext(), &pb.CountDoingHTTPCacheTasksRequest{})
	if err != nil {
		return err
	}

	parent.Data["countDoingTasks"] = countTasksResp.Count
	return nil
}
