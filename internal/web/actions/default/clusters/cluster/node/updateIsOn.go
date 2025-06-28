// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package node

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type UpdateIsOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateIsOnAction) RunPost(params struct {
	NodeId int64
	IsOn   bool
}) {
	if params.IsOn {
		defer this.CreateLogInfo(codes.Node_LogUpdateNodeOn, params.NodeId)
	} else {
		defer this.CreateLogInfo(codes.Node_LogUpdateNodeOff, params.NodeId)
	}

	_, err := this.RPC().NodeRPC().UpdateNodeIsOn(this.AdminContext(), &pb.UpdateNodeIsOnRequest{
		NodeId: params.NodeId,
		IsOn:   params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
