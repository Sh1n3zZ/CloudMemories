// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ocsp

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type ResetAction struct {
	actionutils.ParentAction
}

func (this *ResetAction) RunPost(params struct {
	CertIds []int64
}) {
	defer this.CreateLogInfo(codes.SSLCert_LogOCSPResetOCSPStatus)

	_, err := this.RPC().SSLCertRPC().ResetSSLCertsWithOCSPError(this.AdminContext(), &pb.ResetSSLCertsWithOCSPErrorRequest{SslCertIds: params.CertIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
