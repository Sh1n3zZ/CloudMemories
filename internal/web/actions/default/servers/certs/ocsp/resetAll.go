// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ocsp

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type ResetAllAction struct {
	actionutils.ParentAction
}

func (this *ResetAllAction) RunPost(params struct{}) {
	defer this.CreateLogInfo(codes.SSLCert_LogOCSPResetAllOCSPStatus)

	_, err := this.RPC().SSLCertRPC().ResetAllSSLCertsWithOCSPError(this.AdminContext(), &pb.ResetAllSSLCertsWithOCSPErrorRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
