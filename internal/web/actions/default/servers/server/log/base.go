// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package log

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type BaseAction struct {
	actionutils.ParentAction
}

func (this *BaseAction) initClusterAccessLogConfig(serverId int64) bool {
	this.Data["clusterAccessLogIsOn"] = true
	var clusterId int64
	serverResp, err := this.RPC().ServerRPC().FindEnabledUserServerBasic(this.AdminContext(), &pb.FindEnabledUserServerBasicRequest{ServerId: serverId})
	if err != nil {
		this.ErrorPage(err)
		return false
	}
	if serverResp.Server == nil {
		this.NotFound("Server", serverId)
		return false
	}
	if serverResp.Server.NodeCluster != nil && serverResp.Server.NodeCluster.Id > 0 {
		clusterId = serverResp.Server.NodeCluster.Id
	}

	if clusterId > 0 {
		globalServerConfigResp, err := this.RPC().NodeClusterRPC().FindNodeClusterGlobalServerConfig(this.AdminContext(), &pb.FindNodeClusterGlobalServerConfigRequest{NodeClusterId: clusterId})
		if err != nil {
			this.ErrorPage(err)
			return false
		}

		if len(globalServerConfigResp.GlobalServerConfigJSON) > 0 {
			var globalServerConfig = serverconfigs.NewGlobalServerConfig()
			err = json.Unmarshal(globalServerConfigResp.GlobalServerConfigJSON, globalServerConfig)
			if err != nil {
				this.ErrorPage(err)
				return false
			}
			this.Data["clusterAccessLogIsOn"] = globalServerConfig.HTTPAccessLog.IsOn
		}
	}
	return true
}
