package cluster

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type InstallStatusAction struct {
	actionutils.ParentAction
}

func (this *InstallStatusAction) RunPost(params struct {
	NodeId int64
}) {
	resp, err := this.RPC().NodeRPC().FindNodeInstallStatus(this.AdminContext(), &pb.FindNodeInstallStatusRequest{NodeId: params.NodeId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if resp.InstallStatus == nil {
		this.Data["status"] = nil
		this.Success()
	}

	this.Data["status"] = maps.Map{
		"isRunning":  resp.InstallStatus.IsRunning,
		"isFinished": resp.InstallStatus.IsFinished,
		"isOk":       resp.InstallStatus.IsOk,
		"error":      resp.InstallStatus.Error,
		"errorCode":  resp.InstallStatus.ErrorCode,
	}

	this.Success()
}
