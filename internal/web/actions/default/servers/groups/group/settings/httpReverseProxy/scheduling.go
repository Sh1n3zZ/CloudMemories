package httpReverseProxy

import (
	"encoding/json"
	"errors"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/schedulingconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/groups/group/servergrouputils"
)

type SchedulingAction struct {
	actionutils.ParentAction
}

func (this *SchedulingAction) Init() {
	this.FirstMenu("scheduling")
}

func (this *SchedulingAction) RunGet(params struct {
	GroupId int64
}) {
	_, err := servergrouputils.InitGroup(this.Parent(), params.GroupId, "httpReverseProxy")
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["family"] = "http"

	reverseProxyResp, err := this.RPC().ServerGroupRPC().FindAndInitServerGroupHTTPReverseProxyConfig(this.AdminContext(), &pb.FindAndInitServerGroupHTTPReverseProxyConfigRequest{ServerGroupId: params.GroupId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	reverseProxy := serverconfigs.NewReverseProxyConfig()
	err = json.Unmarshal(reverseProxyResp.ReverseProxyJSON, reverseProxy)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["reverseProxyId"] = reverseProxy.Id

	schedulingCode := reverseProxy.FindSchedulingConfig().Code
	schedulingMap := schedulingconfigs.FindSchedulingType(schedulingCode)
	if schedulingMap == nil {
		this.ErrorPage(errors.New("invalid scheduling code '" + schedulingCode + "'"))
		return
	}
	this.Data["scheduling"] = schedulingMap

	this.Show()
}
