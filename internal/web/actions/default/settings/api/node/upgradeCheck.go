// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package node

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/configs"
	teaconst "github.com/Sh1n3zZ/CloudMemories/internal/const"
	"github.com/Sh1n3zZ/CloudMemories/internal/rpc"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

// UpgradeCheckAction 检查升级结果
type UpgradeCheckAction struct {
	actionutils.ParentAction
}

func (this *UpgradeCheckAction) Init() {
	this.Nav("", "", "")
}

func (this *UpgradeCheckAction) RunPost(params struct {
	NodeId int64
}) {
	this.Data["isOk"] = false

	nodeResp, err := this.RPC().APINodeRPC().FindEnabledAPINode(this.AdminContext(), &pb.FindEnabledAPINodeRequest{ApiNodeId: params.NodeId})
	if err != nil {
		this.Success()
		return
	}

	var node = nodeResp.ApiNode
	if node == nil || len(node.AccessAddrs) == 0 {
		this.Success()
		return
	}

	apiConfig, err := configs.LoadAPIConfig()
	if err != nil {
		this.Success()
		return
	}

	var newAPIConfig = apiConfig.Clone()
	newAPIConfig.RPCEndpoints = node.AccessAddrs
	rpcClient, err := rpc.NewRPCClient(newAPIConfig, false)
	if err != nil {
		this.Success()
		return
	}

	versionResp, err := rpcClient.APINodeRPC().FindCurrentAPINodeVersion(rpcClient.Context(0), &pb.FindCurrentAPINodeVersionRequest{})
	if err != nil {
		this.Success()
		return
	}

	if versionResp.Version != teaconst.Version {
		this.Success()
		return
	}

	this.Data["isOk"] = true

	this.Success()
}
