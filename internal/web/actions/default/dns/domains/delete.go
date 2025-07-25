package domains

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/utils/numberutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	DomainId int64
}) {
	// 记录日志
	defer this.CreateLogInfo(codes.DNS_LogDeleteDomain, params.DomainId)

	// 检查是否正在使用
	countResp, err := this.RPC().NodeClusterRPC().CountAllEnabledNodeClustersWithDNSDomainId(this.AdminContext(), &pb.CountAllEnabledNodeClustersWithDNSDomainIdRequest{DnsDomainId: params.DomainId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if countResp.Count > 0 {
		this.Fail("当前域名正在被" + numberutils.FormatInt64(countResp.Count) + "个集群所使用，所以不能删除。请修改后再操作。")
	}

	// 执行删除
	_, err = this.RPC().DNSDomainRPC().DeleteDNSDomain(this.AdminContext(), &pb.DeleteDNSDomainRequest{DnsDomainId: params.DomainId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
