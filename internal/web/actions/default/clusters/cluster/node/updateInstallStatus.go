package node

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type UpdateInstallStatusAction struct {
	actionutils.ParentAction
}

func (this *UpdateInstallStatusAction) RunPost(params struct {
	NodeId      int64
	IsInstalled bool
}) {
	// 创建日志
	defer this.CreateLogInfo(codes.Node_LogUpdateNodeInstallationStatus, params.NodeId)

	_, err := this.RPC().NodeRPC().UpdateNodeIsInstalled(this.AdminContext(), &pb.UpdateNodeIsInstalledRequest{
		NodeId:      params.NodeId,
		IsInstalled: params.IsInstalled,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
