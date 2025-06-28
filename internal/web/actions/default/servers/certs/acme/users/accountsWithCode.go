// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package users

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type AccountsWithCodeAction struct {
	actionutils.ParentAction
}

func (this *AccountsWithCodeAction) RunPost(params struct {
	Code string
}) {
	accountsResp, err := this.RPC().ACMEProviderAccountRPC().FindAllACMEProviderAccountsWithProviderCode(this.AdminContext(), &pb.FindAllACMEProviderAccountsWithProviderCodeRequest{AcmeProviderCode: params.Code})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var accountMaps = []maps.Map{}
	for _, account := range accountsResp.AcmeProviderAccounts {
		accountMaps = append(accountMaps, maps.Map{
			"id":   account.Id,
			"name": account.Name,
		})
	}
	this.Data["accounts"] = accountMaps

	this.Success()
}
