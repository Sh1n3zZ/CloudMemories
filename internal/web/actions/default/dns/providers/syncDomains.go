// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package providers

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type SyncDomainsAction struct {
	actionutils.ParentAction
}

func (this *SyncDomainsAction) RunPost(params struct {
	ProviderId int64
}) {
	resp, err := this.RPC().DNSDomainRPC().SyncDNSDomainsFromProvider(this.AdminContext(), &pb.SyncDNSDomainsFromProviderRequest{DnsProviderId: params.ProviderId})
	if err != nil {
		this.Fail("更新域名失败：" + err.Error())
	}

	this.Data["hasChanges"] = resp.HasChanges

	this.Success()
}
