package waf

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type SelectPopupAction struct {
	actionutils.ParentAction
}

func (this *SelectPopupAction) Init() {
	this.FirstMenu("index")
}

func (this *SelectPopupAction) RunGet(params struct{}) {
	countResp, err := this.RPC().HTTPFirewallPolicyRPC().CountAllEnabledHTTPFirewallPolicies(this.AdminContext(), &pb.CountAllEnabledHTTPFirewallPoliciesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	count := countResp.Count
	page := this.NewPage(count)

	listResp, err := this.RPC().HTTPFirewallPolicyRPC().ListEnabledHTTPFirewallPolicies(this.AdminContext(), &pb.ListEnabledHTTPFirewallPoliciesRequest{
		Offset: page.Offset,
		Size:   page.Size,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	policyMaps := []maps.Map{}
	for _, policy := range listResp.HttpFirewallPolicies {
		countInbound := 0
		countOutbound := 0
		if len(policy.InboundJSON) > 0 {
			inboundConfig := &firewallconfigs.HTTPFirewallInboundConfig{}
			err = json.Unmarshal(policy.InboundJSON, inboundConfig)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			countInbound = len(inboundConfig.GroupRefs)
		}
		if len(policy.OutboundJSON) > 0 {
			outboundConfig := &firewallconfigs.HTTPFirewallInboundConfig{}
			err = json.Unmarshal(policy.OutboundJSON, outboundConfig)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			countOutbound = len(outboundConfig.GroupRefs)
		}

		policyMaps = append(policyMaps, maps.Map{
			"id":            policy.Id,
			"isOn":          policy.IsOn,
			"name":          policy.Name,
			"countInbound":  countInbound,
			"countOutbound": countOutbound,
		})
	}

	this.Data["policies"] = policyMaps

	this.Data["page"] = page.AsHTML()

	this.Show()
}
