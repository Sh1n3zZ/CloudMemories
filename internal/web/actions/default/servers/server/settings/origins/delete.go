package origins

import (
	"encoding/json"
	"errors"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ReverseProxyId int64
	OriginId       int64
	OriginType     string
}) {
	reverseProxyResp, err := this.RPC().ReverseProxyRPC().FindEnabledReverseProxy(this.AdminContext(), &pb.FindEnabledReverseProxyRequest{ReverseProxyId: params.ReverseProxyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	reverseProxy := reverseProxyResp.ReverseProxy
	if reverseProxy == nil {
		this.ErrorPage(errors.New("reverse proxy is nil"))
		return
	}

	origins := []*serverconfigs.OriginRef{}
	switch params.OriginType {
	case "primary":
		err = json.Unmarshal(reverseProxy.PrimaryOriginsJSON, &origins)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	case "backup":
		err = json.Unmarshal(reverseProxy.BackupOriginsJSON, &origins)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	default:
		this.ErrorPage(errors.New("invalid origin type '" + params.OriginType + "'"))
		return
	}

	result := []*serverconfigs.OriginRef{}
	for _, origin := range origins {
		if origin.OriginId == params.OriginId {
			continue
		}
		result = append(result, origin)
	}
	resultData, err := json.Marshal(result)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	switch params.OriginType {
	case "primary":
		_, err = this.RPC().ReverseProxyRPC().UpdateReverseProxyPrimaryOrigins(this.AdminContext(), &pb.UpdateReverseProxyPrimaryOriginsRequest{
			ReverseProxyId: params.ReverseProxyId,
			OriginsJSON:    resultData,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	case "backup":
		_, err = this.RPC().ReverseProxyRPC().UpdateReverseProxyBackupOrigins(this.AdminContext(), &pb.UpdateReverseProxyBackupOriginsRequest{
			ReverseProxyId: params.ReverseProxyId,
			OriginsJSON:    resultData,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	defer this.CreateLogInfo(codes.ServerOrigin_LogDeleteOrigin, params.OriginId)

	this.Success()
}
