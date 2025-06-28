package regions

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type SortAction struct {
	actionutils.ParentAction
}

func (this *SortAction) RunPost(params struct {
	RegionIds []int64
}) {
	defer this.CreateLogInfo(codes.NodeRegion_LogSortNodeRegions)

	_, err := this.RPC().NodeRegionRPC().UpdateNodeRegionOrders(this.AdminContext(), &pb.UpdateNodeRegionOrdersRequest{NodeRegionIds: params.RegionIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
