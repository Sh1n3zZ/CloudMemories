// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ocsp

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type IgnoreAction struct {
	actionutils.ParentAction
}

func (this *IgnoreAction) RunPost(params struct {
	CertIds []int64
}) {
	defer this.CreateLogInfo(codes.SSLCert_LogOCSPIgnoreOCSPStatus)

	_, err := this.RPC().SSLCertRPC().IgnoreSSLCertsWithOCSPError(this.AdminContext(), &pb.IgnoreSSLCertsWithOCSPErrorRequest{SslCertIds: params.CertIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
