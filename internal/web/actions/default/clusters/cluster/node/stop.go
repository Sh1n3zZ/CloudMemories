package node

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type StopAction struct {
	actionutils.ParentAction
}

func (this *StopAction) RunPost(params struct {
	NodeId int64
}) {
	resp, err := this.RPC().NodeRPC().StopNode(this.AdminContext(), &pb.StopNodeRequest{NodeId: params.NodeId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLogInfo(codes.Node_LogStopNodeRemotely, params.NodeId)

	if resp.IsOk {
		this.Success()
	}

	this.Fail("执行失败：" + resp.Error)
}
