package reverseProxy

import (
	"encoding/json"
	"errors"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/schedulingconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type SchedulingAction struct {
	actionutils.ParentAction
}

func (this *SchedulingAction) Init() {
	this.FirstMenu("scheduling")
}

func (this *SchedulingAction) RunGet(params struct {
	ServerId int64
}) {
	reverseProxyResp, err := this.RPC().ServerRPC().FindAndInitServerReverseProxyConfig(this.AdminContext(), &pb.FindAndInitServerReverseProxyConfigRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var reverseProxy = serverconfigs.NewReverseProxyConfig()
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
