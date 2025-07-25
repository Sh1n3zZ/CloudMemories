package serverutils

import (
	"errors"
	"strconv"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

// FindServer 查找服务信息
func FindServer(p *actionutils.ParentAction, serverId int64) (*pb.Server, *serverconfigs.ServerConfig, bool) {
	serverResp, err := p.RPC().ServerRPC().FindEnabledServer(p.AdminContext(), &pb.FindEnabledServerRequest{
		ServerId:       serverId,
		IgnoreSSLCerts: true,
	})
	if err != nil {
		p.ErrorPage(err)
		return nil, nil, false
	}
	var server = serverResp.Server
	if server == nil {
		p.ErrorPage(errors.New("not found server with id '" + strconv.FormatInt(serverId, 10) + "'"))
		return nil, nil, false
	}
	config, err := serverconfigs.NewServerConfigFromJSON(server.Config)
	if err != nil {
		p.ErrorPage(err)
		return nil, nil, false
	}

	return server, config, true
}
