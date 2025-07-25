package waf

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/dao"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type GroupsAction struct {
	actionutils.ParentAction
}

func (this *GroupsAction) Init() {
	this.Nav("", "setting", this.ParamString("type"))
	this.SecondMenu("waf")
}

func (this *GroupsAction) RunGet(params struct {
	ServerId         int64
	FirewallPolicyId int64
	Type             string
}) {
	this.Data["firewallPolicyId"] = params.FirewallPolicyId
	this.Data["type"] = params.Type

	firewallPolicy, err := dao.SharedHTTPFirewallPolicyDAO.FindEnabledHTTPFirewallPolicyConfig(this.AdminContext(), params.FirewallPolicyId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if firewallPolicy == nil {
		this.NotFound("firewallPolicy", params.FirewallPolicyId)
		return
	}

	groupMaps := []maps.Map{}

	// inbound
	if params.Type == "inbound" {
		if firewallPolicy.Inbound != nil {
			for _, g := range firewallPolicy.Inbound.Groups {
				groupMaps = append(groupMaps, maps.Map{
					"id":          g.Id,
					"name":        g.Name,
					"code":        g.Code,
					"isOn":        g.IsOn,
					"description": g.Description,
					"countSets":   len(g.Sets),
					"canDelete":   len(g.Code) == 0,
				})
			}
		}
	}

	// outbound
	if params.Type == "outbound" {
		if firewallPolicy.Outbound != nil {
			for _, g := range firewallPolicy.Outbound.Groups {
				groupMaps = append(groupMaps, maps.Map{
					"id":          g.Id,
					"name":        g.Name,
					"code":        g.Code,
					"isOn":        g.IsOn,
					"description": g.Description,
					"countSets":   len(g.Sets),
					"canDelete":   len(g.Code) == 0,
				})
			}
		}
	}

	this.Data["groups"] = groupMaps

	// WAF是否启用
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerId(this.AdminContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["wafIsOn"] = webConfig.FirewallRef != nil && webConfig.FirewallRef.IsOn

	this.Show()
}
