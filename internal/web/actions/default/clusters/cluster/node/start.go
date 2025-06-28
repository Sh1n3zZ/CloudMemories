package node

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type StartAction struct {
	actionutils.ParentAction
}

func (this *StartAction) RunPost(params struct {
	NodeId int64
}) {
	resp, err := this.RPC().NodeRPC().StartNode(this.AdminContext(), &pb.StartNodeRequest{NodeId: params.NodeId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLogInfo(codes.Node_LogStartNodeRemotely, params.NodeId)

	if resp.IsOk {
		this.Success()
	}

	this.Fail("启动失败：" + resp.Error)
}
