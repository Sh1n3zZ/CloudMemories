// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package logs

import (
	"strings"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

// PartitionDataAction 读取分区表
type PartitionDataAction struct {
	actionutils.ParentAction
}

func (this *PartitionDataAction) RunPost(params struct {
	Day string
}) {
	var day = params.Day
	day = strings.ReplaceAll(day, "-", "")

	resp, err := this.RPC().HTTPAccessLogRPC().FindHTTPAccessLogPartitions(this.AdminContext(), &pb.FindHTTPAccessLogPartitionsRequest{
		Day: day,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if len(resp.Partitions) > 0 {
		this.Data["partitions"] = resp.Partitions
	} else {
		this.Data["partitions"] = []int32{}
	}

	this.Success()
}
