// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dns

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type RegenerateCNAMEAction struct {
	actionutils.ParentAction
}

func (this *RegenerateCNAMEAction) RunPost(params struct {
	ServerId int64
}) {
	defer this.CreateLogInfo(codes.ServerDNS_LogRegenerateDNSName, params.ServerId)

	_, err := this.RPC().ServerRPC().RegenerateServerDNSName(this.AdminContext(), &pb.RegenerateServerDNSNameRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
