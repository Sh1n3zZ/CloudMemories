package groups

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type SortAction struct {
	actionutils.ParentAction
}

func (this *SortAction) RunPost(params struct {
	GroupIds []int64
}) {
	_, err := this.RPC().NodeGroupRPC().UpdateNodeGroupOrders(this.AdminContext(), &pb.UpdateNodeGroupOrdersRequest{NodeGroupIds: params.GroupIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLogInfo(codes.NodeGroup_LogSortNodeGroups)

	this.Success()
}
