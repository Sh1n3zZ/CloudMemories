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
	// 创建日志
	defer this.CreateLogInfo(codes.ServerGroup_LogSortServerGroups)

	_, err := this.RPC().ServerGroupRPC().UpdateServerGroupOrders(this.AdminContext(), &pb.UpdateServerGroupOrdersRequest{ServerGroupIds: params.GroupIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
