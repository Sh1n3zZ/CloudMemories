// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package metrics

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ClusterId int64
	ItemId    int64
}) {
	defer this.CreateLogInfo(codes.MetricItem_LogDeleteMetricItemFromCluster, params.ClusterId, params.ItemId)

	_, err := this.RPC().NodeClusterMetricItemRPC().DisableNodeClusterMetricItem(this.AdminContext(), &pb.DisableNodeClusterMetricItemRequest{
		NodeClusterId: params.ClusterId,
		MetricItemId:  params.ItemId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
