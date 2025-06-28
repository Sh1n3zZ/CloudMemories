package domains

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type SyncAction struct {
	actionutils.ParentAction
}

func (this *SyncAction) RunPost(params struct {
	DomainId int64
}) {
	// 记录日志
	defer this.CreateLogInfo(codes.DNS_LogSyncDomain, params.DomainId)

	// 执行同步
	resp, err := this.RPC().DNSDomainRPC().SyncDNSDomainData(this.AdminContext(), &pb.SyncDNSDomainDataRequest{DnsDomainId: params.DomainId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if resp.IsOk {
		this.Success()
	} else {
		this.Data["shouldFix"] = resp.ShouldFix
		this.Fail(resp.Error)
	}

	this.Success()
}
