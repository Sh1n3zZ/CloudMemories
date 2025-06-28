package server

import (
	"strconv"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	teaconst "github.com/Sh1n3zZ/CloudMemories/internal/const"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "index", "index")
	this.SecondMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	// TODO 等看板实现后，需要跳转到看板

	// TCP & UDP跳转到设置
	serverTypeResp, err := this.RPC().ServerRPC().FindEnabledServerType(this.AdminContext(), &pb.FindEnabledServerTypeRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverType := serverTypeResp.Type
	if serverType == serverconfigs.ServerTypeTCPProxy || serverType == serverconfigs.ServerTypeUDPProxy {
		this.RedirectURL("/servers/server/settings?serverId=" + strconv.FormatInt(params.ServerId, 10))
		return
	}

	// HTTP跳转到访问日志
	if teaconst.IsPlus {
		this.RedirectURL("/servers/server/boards?serverId=" + strconv.FormatInt(params.ServerId, 10))
	} else {
		this.RedirectURL("/servers/server/stat?serverId=" + strconv.FormatInt(params.ServerId, 10))
	}
}
