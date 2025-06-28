// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package accounts

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	AccountId int64
}) {
	defer this.CreateLogInfo(codes.ACMEProviderAccount_LogDeleteACMEProviderAccount, params.AccountId)

	_, err := this.RPC().ACMEProviderAccountRPC().DeleteACMEProviderAccount(this.AdminContext(), &pb.DeleteACMEProviderAccountRequest{AcmeProviderAccountId: params.AccountId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
