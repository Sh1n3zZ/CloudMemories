package nodes

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
	NodeId    int64
}) {
	// 创建日志
	defer this.CreateLogInfo(codes.Node_LogDeleteNodeFromCluster, params.ClusterId, params.NodeId)

	_, err := this.RPC().NodeRPC().DeleteNodeFromNodeCluster(this.AdminContext(), &pb.DeleteNodeFromNodeClusterRequest{
		NodeId:        params.NodeId,
		NodeClusterId: params.ClusterId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
