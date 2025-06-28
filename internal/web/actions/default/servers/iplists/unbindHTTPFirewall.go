// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package iplists

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/dao"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type UnbindHTTPFirewallAction struct {
	actionutils.ParentAction
}

func (this *UnbindHTTPFirewallAction) RunPost(params struct {
	HttpFirewallPolicyId int64
	ListId               int64
}) {
	defer this.CreateLogInfo(codes.IPList_LogUnbindIPListWAFPolicy, params.ListId, params.HttpFirewallPolicyId)

	// List类型
	listResp, err := this.RPC().IPListRPC().FindEnabledIPList(this.AdminContext(), &pb.FindEnabledIPListRequest{IpListId: params.ListId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var list = listResp.IpList
	if list == nil {
		this.Fail("找不到要使用的IP名单")
	}

	// 已经绑定的
	inboundConfig, err := dao.SharedHTTPFirewallPolicyDAO.FindEnabledHTTPFirewallPolicyInboundConfig(this.AdminContext(), params.HttpFirewallPolicyId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if inboundConfig == nil {
		inboundConfig = &firewallconfigs.HTTPFirewallInboundConfig{IsOn: true}
	}
	inboundConfig.RemovePublicList(list.Id, list.Type)

	inboundJSON, err := json.Marshal(inboundConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallInboundConfig(this.AdminContext(), &pb.UpdateHTTPFirewallInboundConfigRequest{
		HttpFirewallPolicyId: params.HttpFirewallPolicyId,
		InboundJSON:          inboundJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
