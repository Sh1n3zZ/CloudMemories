package cluster

import (
	"strings"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type InstallNodesAction struct {
	actionutils.ParentAction
}

func (this *InstallNodesAction) Init() {
	this.Nav("", "node", "install")
	this.SecondMenu("nodes")
}

func (this *InstallNodesAction) RunGet(params struct {
	ClusterId int64
}) {
	this.Data["leftMenuItems"] = LeftMenuItemsForInstall(this.AdminContext(), params.ClusterId, "register", this.LangCode())

	clusterResp, err := this.RPC().NodeClusterRPC().FindEnabledNodeCluster(this.AdminContext(), &pb.FindEnabledNodeClusterRequest{NodeClusterId: params.ClusterId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if clusterResp.NodeCluster == nil {
		this.NotFound("nodeCluster", params.ClusterId)
		return
	}

	cluster := clusterResp.NodeCluster

	clusterAPINodesResp, err := this.RPC().NodeClusterRPC().FindAPINodesWithNodeCluster(this.AdminContext(), &pb.FindAPINodesWithNodeClusterRequest{NodeClusterId: params.ClusterId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if err != nil {
		this.ErrorPage(err)
		return
	}
	apiNodeAddrs := []string{}
	if clusterAPINodesResp.UseAllAPINodes {
		apiNodesResp, err := this.RPC().APINodeRPC().FindAllEnabledAPINodes(this.AdminContext(), &pb.FindAllEnabledAPINodesRequest{})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		for _, apiNode := range apiNodesResp.ApiNodes {
			if !apiNode.IsOn {
				continue
			}
			apiNodeAddrs = append(apiNodeAddrs, apiNode.AccessAddrs...)
		}
	} else {
		for _, apiNode := range clusterAPINodesResp.ApiNodes {
			if !apiNode.IsOn {
				continue
			}
			apiNodeAddrs = append(apiNodeAddrs, apiNode.AccessAddrs...)
		}
	}

	this.Data["cluster"] = maps.Map{
		"uniqueId":  cluster.UniqueId,
		"secret":    cluster.Secret,
		"endpoints": "\"" + strings.Join(apiNodeAddrs, "\", \"") + "\"",
	}

	this.Show()
}
