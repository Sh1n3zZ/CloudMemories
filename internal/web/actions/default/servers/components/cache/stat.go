package cache

import (
	"net/http"
	"strconv"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/messageconfigs"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/nodes/nodeutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
)

type StatAction struct {
	actionutils.ParentAction
}

func (this *StatAction) Init() {
	this.Nav("", "", "stat")
}

func (this *StatAction) RunGet(params struct {
	CachePolicyId int64
}) {
	// 默认的集群ID
	cookie, err := this.Request.Cookie("cache_cluster_id")
	if cookie != nil && err == nil {
		this.Data["clusterId"] = types.Int64(cookie.Value)
	}

	// 集群列表
	clustersResp, err := this.RPC().NodeClusterRPC().FindAllEnabledNodeClustersWithHTTPCachePolicyId(this.AdminContext(), &pb.FindAllEnabledNodeClustersWithHTTPCachePolicyIdRequest{HttpCachePolicyId: params.CachePolicyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	clusterMaps := []maps.Map{}
	for _, cluster := range clustersResp.NodeClusters {
		clusterMaps = append(clusterMaps, maps.Map{
			"id":   cluster.Id,
			"name": cluster.Name,
		})
	}
	this.Data["clusters"] = clusterMaps

	this.Show()
}

func (this *StatAction) RunPost(params struct {
	CachePolicyId int64
	ClusterId     int64

	Must *actions.Must
}) {
	// 记录clusterId
	this.AddCookie(&http.Cookie{
		Name:  "cache_cluster_id",
		Value: strconv.FormatInt(params.ClusterId, 10),
	})

	cachePolicyResp, err := this.RPC().HTTPCachePolicyRPC().FindEnabledHTTPCachePolicyConfig(this.AdminContext(), &pb.FindEnabledHTTPCachePolicyConfigRequest{HttpCachePolicyId: params.CachePolicyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	cachePolicyJSON := cachePolicyResp.HttpCachePolicyJSON
	if len(cachePolicyJSON) == 0 {
		this.Fail("找不到要操作的缓存策略")
	}

	// 发送命令
	msg := &messageconfigs.StatCacheMessage{
		CachePolicyJSON: cachePolicyJSON,
	}
	results, err := nodeutils.SendMessageToCluster(this.AdminContext(), params.ClusterId, messageconfigs.MessageCodeStatCache, msg, 10)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	isAllOk := true
	for _, result := range results {
		if !result.IsOK {
			isAllOk = false
			break
		}
	}

	this.Data["isAllOk"] = isAllOk
	this.Data["results"] = results

	// 创建日志
	defer this.CreateLogInfo(codes.ServerCachePolicy_LogStatCaches, params.CachePolicyId)

	this.Success()
}
