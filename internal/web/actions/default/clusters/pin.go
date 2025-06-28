// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package clusters

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type PinAction struct {
	actionutils.ParentAction
}

func (this *PinAction) RunPost(params struct {
	ClusterId int64
	IsPinned  bool
}) {
	if params.IsPinned {
		defer this.CreateLogInfo(codes.NodeCluster_LogPinCluster, params.ClusterId)
	} else {
		defer this.CreateLogInfo(codes.NodeCluster_LogUnpinCluster, params.ClusterId)
	}

	_, err := this.RPC().NodeClusterRPC().UpdateNodeClusterPinned(this.AdminContext(), &pb.UpdateNodeClusterPinnedRequest{
		NodeClusterId: params.ClusterId,
		IsPinned:      params.IsPinned,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
