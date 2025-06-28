// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package clusters

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type NodeOptionsAction struct {
	actionutils.ParentAction
}

func (this *NodeOptionsAction) RunPost(params struct {
	ClusterId int64
}) {
	resp, err := this.RPC().NodeRPC().FindAllEnabledNodesWithNodeClusterId(this.AdminContext(), &pb.FindAllEnabledNodesWithNodeClusterIdRequest{NodeClusterId: params.ClusterId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	var nodeMaps = []maps.Map{}
	for _, node := range resp.Nodes {
		nodeMaps = append(nodeMaps, maps.Map{
			"id":   node.Id,
			"name": node.Name,
		})
	}
	this.Data["nodes"] = nodeMaps

	this.Success()
}
