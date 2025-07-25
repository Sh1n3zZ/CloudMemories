package firewallActions

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type UpdatePopupAction struct {
	actionutils.ParentAction
}

func (this *UpdatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdatePopupAction) RunGet(params struct {
	ActionId int64
}) {
	actionResp, err := this.RPC().NodeClusterFirewallActionRPC().FindEnabledNodeClusterFirewallAction(this.AdminContext(), &pb.FindEnabledNodeClusterFirewallActionRequest{NodeClusterFirewallActionId: params.ActionId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	action := actionResp.NodeClusterFirewallAction
	if action == nil {
		this.NotFound("nodeClusterFirewallAction", params.ActionId)
		return
	}

	actionParams := maps.Map{}
	if len(action.ParamsJSON) > 0 {
		err = json.Unmarshal(action.ParamsJSON, &actionParams)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	this.Data["action"] = maps.Map{
		"id":         action.Id,
		"name":       action.Name,
		"eventLevel": action.EventLevel,
		"params":     actionParams,
		"type":       action.Type,
	}

	// 通用参数
	this.Data["actionTypes"] = firewallconfigs.FindAllFirewallActionTypes()

	this.Show()
}

func (this *UpdatePopupAction) RunPost(params struct {
	ActionId   int64
	Name       string
	EventLevel string
	Type       string

	// ipset
	IpsetWhiteName          string
	IpsetBlackName          string
	IpsetWhiteNameIPv6      string
	IpsetBlackNameIPv6      string
	IpsetAutoAddToIPTables  bool
	IpsetAutoAddToFirewalld bool

	// script
	ScriptPath string

	// http api
	HttpAPIURL string

	// HTML内容
	HtmlContent string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo(codes.WAFAction_LogUpdateWAFAction, params.ActionId)

	params.Must.
		Field("name", params.Name).
		Require("请输入动作名称").
		Field("type", params.Type).
		Require("请选择动作类型")

	var actionParams interface{} = nil
	switch params.Type {
	case firewallconfigs.FirewallActionTypeIPSet:
		params.Must.
			Field("ipsetWhiteName", params.IpsetWhiteName).
			Require("请输入IPSet白名单名称").
			Match(`^\w+$`, "请输入正确的IPSet白名单名称").
			Field("ipsetBlackName", params.IpsetBlackName).
			Require("请输入IPSet黑名单名称").
			Match(`^\w+$`, "请输入正确的IPSet黑名单名称").
			Field("ipsetWhiteNameIPv6", params.IpsetWhiteNameIPv6).
			Require("请输入IPSet IPv6白名单名称").
			Match(`^\w+$`, "请输入正确的IPSet IPv6白名单名称").
			Field("ipsetBlackNameIPv6", params.IpsetBlackNameIPv6).
			Require("请输入IPSet IPv6黑名单名称").
			Match(`^\w+$`, "请输入正确的IPSet IPv6黑名单名称")

		actionParams = &firewallconfigs.FirewallActionIPSetConfig{
			WhiteName:          params.IpsetWhiteName,
			BlackName:          params.IpsetBlackName,
			WhiteNameIPv6:      params.IpsetWhiteNameIPv6,
			BlackNameIPv6:      params.IpsetBlackNameIPv6,
			AutoAddToIPTables:  params.IpsetAutoAddToIPTables,
			AutoAddToFirewalld: params.IpsetAutoAddToFirewalld,
		}
	case firewallconfigs.FirewallActionTypeIPTables:
		actionParams = &firewallconfigs.FirewallActionIPTablesConfig{}
	case firewallconfigs.FirewallActionTypeFirewalld:
		actionParams = &firewallconfigs.FirewallActionFirewalldConfig{}
	case firewallconfigs.FirewallActionTypeScript:
		params.Must.
			Field("scriptPath", params.ScriptPath).
			Require("请输入脚本路径")
		actionParams = &firewallconfigs.FirewallActionScriptConfig{
			Path: params.ScriptPath,
		}
	case firewallconfigs.FirewallActionTypeHTTPAPI:
		params.Must.
			Field("httpAPIURL", params.HttpAPIURL).
			Require("请输入API URL").
			Match(`^(http|https):`, "API地址必须以http://或https://开头")
		actionParams = &firewallconfigs.FirewallActionHTTPAPIConfig{
			URL: params.HttpAPIURL,
		}
	case firewallconfigs.FirewallActionTypeHTML:
		params.Must.
			Field("htmlContent", params.HtmlContent).
			Require("请输入HTML内容")
		actionParams = &firewallconfigs.FirewallActionHTMLConfig{
			Content: params.HtmlContent,
		}
	default:
		this.Fail("选择的类型'" + params.Type + "'暂时不支持")
	}

	actionParamsJSON, err := json.Marshal(actionParams)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = this.RPC().NodeClusterFirewallActionRPC().UpdateNodeClusterFirewallAction(this.AdminContext(), &pb.UpdateNodeClusterFirewallActionRequest{
		NodeClusterFirewallActionId: params.ActionId,
		Name:                        params.Name,
		EventLevel:                  params.EventLevel,
		Type:                        params.Type,
		ParamsJSON:                  actionParamsJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
