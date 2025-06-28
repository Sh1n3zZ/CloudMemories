package waf

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type UpdateSetOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateSetOnAction) RunPost(params struct {
	SetId int64
	IsOn  bool
}) {
	// 日志
	defer this.CreateLogInfo(codes.WAFRuleSet_LogUpdateRuleSetIsOn, params.SetId)

	_, err := this.RPC().HTTPFirewallRuleSetRPC().UpdateHTTPFirewallRuleSetIsOn(this.AdminContext(), &pb.UpdateHTTPFirewallRuleSetIsOnRequest{
		FirewallRuleSetId: params.SetId,
		IsOn:              params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
