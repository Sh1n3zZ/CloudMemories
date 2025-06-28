package db

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	NodeId int64
}) {
	// 创建日志
	defer this.CreateLogInfo(codes.DBNode_LogDeleteDBNode, params.NodeId)

	_, err := this.RPC().DBNodeRPC().DeleteDBNode(this.AdminContext(), &pb.DeleteDBNodeRequest{DbNodeId: params.NodeId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
