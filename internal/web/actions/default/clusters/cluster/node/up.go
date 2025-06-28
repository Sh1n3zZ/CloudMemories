package node

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

// 手动上线
type UpAction struct {
	actionutils.ParentAction
}

func (this *UpAction) RunPost(params struct {
	NodeId int64
}) {
	defer this.CreateLogInfo(codes.Node_LogUpNode, params.NodeId)

	_, err := this.RPC().NodeRPC().UpdateNodeUp(this.AdminContext(), &pb.UpdateNodeUpRequest{
		NodeId: params.NodeId,
		IsUp:   true,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
