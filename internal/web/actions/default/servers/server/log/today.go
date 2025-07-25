package log

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/iplibrary"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/dao"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type TodayAction struct {
	BaseAction
}

func (this *TodayAction) Init() {
	this.Nav("", "log", "")
	this.SecondMenu("today")
}

func (this *TodayAction) RunGet(params struct {
	RequestId string
	ServerId  int64
	HasError  int
	HasWAF    int
	Keyword   string
	Ip        string
	Domain    string
	ClusterId int64
	NodeId    int64

	Partition int32 `default:"-1"`

	PageSize int
}) {
	this.Data["pageSize"] = params.PageSize

	var size = int64(params.PageSize)
	if size < 1 {
		size = 20
	}

	this.Data["path"] = this.Request.URL.Path
	this.Data["hasError"] = params.HasError
	this.Data["keyword"] = params.Keyword
	this.Data["ip"] = params.Ip
	this.Data["domain"] = params.Domain
	this.Data["hasWAF"] = params.HasWAF
	this.Data["clusterId"] = params.ClusterId
	this.Data["nodeId"] = params.NodeId
	this.Data["partition"] = params.Partition
	this.Data["day"] = timeutil.Format("Ymd")

	// 检查集群全局设置
	if !this.initClusterAccessLogConfig(params.ServerId) {
		return
	}

	// 检查当前网站有无开启访问日志
	this.Data["serverAccessLogIsOn"] = true

	groupResp, err := this.RPC().ServerGroupRPC().FindEnabledServerGroupConfigInfo(this.AdminContext(), &pb.FindEnabledServerGroupConfigInfoRequest{
		ServerId: params.ServerId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if !groupResp.HasAccessLogConfig {
		webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerId(this.AdminContext(), params.ServerId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if webConfig != nil && webConfig.AccessLogRef != nil && !webConfig.AccessLogRef.IsOn {
			this.Data["serverAccessLogIsOn"] = false
		}
	}

	resp, err := this.RPC().HTTPAccessLogRPC().ListHTTPAccessLogs(this.AdminContext(), &pb.ListHTTPAccessLogsRequest{
		Partition:         params.Partition,
		RequestId:         params.RequestId,
		ServerId:          params.ServerId,
		HasError:          params.HasError > 0,
		HasFirewallPolicy: params.HasWAF > 0,
		Day:               timeutil.Format("Ymd"),
		Keyword:           params.Keyword,
		Ip:                params.Ip,
		Domain:            params.Domain,
		NodeId:            params.NodeId,
		NodeClusterId:     params.ClusterId,
		Size:              size,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	var ipList = []string{}
	var wafMaps = []maps.Map{}

	if len(resp.HttpAccessLogs) == 0 {
		this.Data["accessLogs"] = []interface{}{}
	} else {
		this.Data["accessLogs"] = resp.HttpAccessLogs
		for _, accessLog := range resp.HttpAccessLogs {
			// IP
			if len(accessLog.RemoteAddr) > 0 {
				if !lists.ContainsString(ipList, accessLog.RemoteAddr) {
					ipList = append(ipList, accessLog.RemoteAddr)
				}
			}

			// WAF信息集合
			if accessLog.FirewallPolicyId > 0 && accessLog.FirewallRuleGroupId > 0 && accessLog.FirewallRuleSetId > 0 {
				// 检查Set是否已经存在
				var existSet = false
				for _, wafMap := range wafMaps {
					if wafMap.GetInt64("setId") == accessLog.FirewallRuleSetId {
						existSet = true
						break
					}
				}
				if !existSet {
					wafMaps = append(wafMaps, maps.Map{
						"policyId": accessLog.FirewallPolicyId,
						"groupId":  accessLog.FirewallRuleGroupId,
						"setId":    accessLog.FirewallRuleSetId,
					})
				}
			}
		}
	}
	this.Data["hasMore"] = resp.HasMore
	this.Data["nextRequestId"] = resp.RequestId

	// 上一个requestId
	this.Data["hasPrev"] = false
	this.Data["lastRequestId"] = ""
	if len(params.RequestId) > 0 {
		this.Data["hasPrev"] = true
		prevResp, err := this.RPC().HTTPAccessLogRPC().ListHTTPAccessLogs(this.AdminContext(), &pb.ListHTTPAccessLogsRequest{
			Partition:         params.Partition,
			RequestId:         params.RequestId,
			ServerId:          params.ServerId,
			HasError:          params.HasError > 0,
			HasFirewallPolicy: params.HasWAF > 0,
			Day:               timeutil.Format("Ymd"),
			Keyword:           params.Keyword,
			Ip:                params.Ip,
			Domain:            params.Domain,
			NodeId:            params.NodeId,
			NodeClusterId:     params.ClusterId,
			Size:              size,
			Reverse:           true,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if int64(len(prevResp.HttpAccessLogs)) == size {
			this.Data["lastRequestId"] = prevResp.RequestId
		}
	}

	// 根据IP查询区域
	this.Data["regions"] = iplibrary.LookupIPSummaries(ipList)

	// WAF相关
	var wafInfos = map[int64]maps.Map{}                          // set id => WAF Map
	var wafPolicyCacheMap = map[int64]*pb.HTTPFirewallPolicy{}   // id => *pb.HTTPFirewallPolicy
	var wafGroupCacheMap = map[int64]*pb.HTTPFirewallRuleGroup{} // id => *pb.HTTPFirewallRuleGroup
	var wafSetCacheMap = map[int64]*pb.HTTPFirewallRuleSet{}     // id => *pb.HTTPFirewallRuleSet
	for _, wafMap := range wafMaps {
		var policyId = wafMap.GetInt64("policyId")
		var groupId = wafMap.GetInt64("groupId")
		var setId = wafMap.GetInt64("setId")
		if policyId > 0 {
			pbPolicy, ok := wafPolicyCacheMap[policyId]
			if !ok {
				policyResp, err := this.RPC().HTTPFirewallPolicyRPC().FindEnabledHTTPFirewallPolicy(this.AdminContext(), &pb.FindEnabledHTTPFirewallPolicyRequest{HttpFirewallPolicyId: policyId})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				pbPolicy = policyResp.HttpFirewallPolicy
				wafPolicyCacheMap[policyId] = pbPolicy
			}
			if pbPolicy != nil {
				wafMap = maps.Map{
					"policy": maps.Map{
						"id":       pbPolicy.Id,
						"name":     pbPolicy.Name,
						"serverId": pbPolicy.ServerId,
					},
				}
				if groupId > 0 {
					pbGroup, ok := wafGroupCacheMap[groupId]
					if !ok {
						groupResp, err := this.RPC().HTTPFirewallRuleGroupRPC().FindEnabledHTTPFirewallRuleGroup(this.AdminContext(), &pb.FindEnabledHTTPFirewallRuleGroupRequest{FirewallRuleGroupId: groupId})
						if err != nil {
							this.ErrorPage(err)
							return
						}
						pbGroup = groupResp.FirewallRuleGroup
						wafGroupCacheMap[groupId] = pbGroup
					}

					if pbGroup != nil {
						wafMap["group"] = maps.Map{
							"id":   pbGroup.Id,
							"name": pbGroup.Name,
						}

						if setId > 0 {
							pbSet, ok := wafSetCacheMap[setId]
							if !ok {
								setResp, err := this.RPC().HTTPFirewallRuleSetRPC().FindEnabledHTTPFirewallRuleSet(this.AdminContext(), &pb.FindEnabledHTTPFirewallRuleSetRequest{FirewallRuleSetId: setId})
								if err != nil {
									this.ErrorPage(err)
									return
								}
								pbSet = setResp.FirewallRuleSet
								wafSetCacheMap[setId] = pbSet
							}

							if pbSet != nil {
								wafMap["set"] = maps.Map{
									"id":   pbSet.Id,
									"name": pbSet.Name,
								}
							}
						}
					}
				}
			}
		}

		wafInfos[setId] = wafMap
	}
	this.Data["wafInfos"] = wafInfos

	this.Show()
}
