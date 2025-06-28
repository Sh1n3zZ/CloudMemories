package cluster

import (
	"context"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs"
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/rpc"
	"github.com/Sh1n3zZ/CloudMemories/internal/utils/numberutils"
	"github.com/iwind/TeaGo/maps"
)

// LeftMenuItemsForInstall 安装升级相关的左侧菜单
func LeftMenuItemsForInstall(ctx context.Context, clusterId int64, selectedItem string, langCode string) []maps.Map {
	rpcClient, _ := rpc.SharedRPC()
	countNotInstalled := int64(0)
	countUpgrade := int64(0)
	if rpcClient != nil {
		{
			resp, err := rpcClient.NodeRPC().CountAllNotInstalledNodesWithNodeClusterId(ctx, &pb.CountAllNotInstalledNodesWithNodeClusterIdRequest{NodeClusterId: clusterId})
			if err == nil {
				countNotInstalled = resp.Count
			}
		}
		{
			resp, err := rpcClient.NodeRPC().CountAllUpgradeNodesWithNodeClusterId(ctx, &pb.CountAllUpgradeNodesWithNodeClusterIdRequest{NodeClusterId: clusterId})
			if err == nil {
				countUpgrade = resp.Count
			}
		}
	}

	return []maps.Map{
		{
			"name":     langs.Message(langCode, codes.NodeMenu_InstallManually),
			"url":      "/clusters/cluster/installManual?clusterId=" + numberutils.FormatInt64(clusterId),
			"isActive": selectedItem == "manual",
		},
		{
			"name":     langs.Message(langCode, codes.NodeMenu_InstallAutoRegister),
			"url":      "/clusters/cluster/installNodes?clusterId=" + numberutils.FormatInt64(clusterId),
			"isActive": selectedItem == "register",
		},
		{
			"name":     langs.Message(langCode, codes.NodeMenu_InstallRemote, countNotInstalled),
			"url":      "/clusters/cluster/installRemote?clusterId=" + numberutils.FormatInt64(clusterId),
			"isActive": selectedItem == "install",
		},
		{
			"name":     langs.Message(langCode, codes.NodeMenu_InstallRemoteUpgrade, countUpgrade),
			"url":      "/clusters/cluster/upgradeRemote?clusterId=" + numberutils.FormatInt64(clusterId),
			"isActive": selectedItem == "upgrade",
		},
	}
}
