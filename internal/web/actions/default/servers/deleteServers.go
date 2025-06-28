// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package servers

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

// DeleteServersAction 删除一组网站
type DeleteServersAction struct {
	actionutils.ParentAction
}

func (this *DeleteServersAction) RunPost(params struct {
	ServerIds []int64
}) {
	defer this.CreateLogInfo(codes.Server_LogDeleteServers)

	_, err := this.RPC().ServerRPC().DeleteServers(this.AdminContext(), &pb.DeleteServersRequest{ServerIds: params.ServerIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
