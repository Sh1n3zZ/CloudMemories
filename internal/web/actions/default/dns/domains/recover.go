package domains

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type RecoverAction struct {
	actionutils.ParentAction
}

func (this *RecoverAction) RunPost(params struct {
	DomainId int64
}) {
	// 记录日志
	defer this.CreateLogInfo(codes.DNS_LogRecoverDomain, params.DomainId)

	// 执行恢复
	_, err := this.RPC().DNSDomainRPC().RecoverDNSDomain(this.AdminContext(), &pb.RecoverDNSDomainRequest{DnsDomainId: params.DomainId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
