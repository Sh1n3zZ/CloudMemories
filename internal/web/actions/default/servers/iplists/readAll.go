// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package iplists

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
)

type ReadAllAction struct {
	actionutils.ParentAction
}

func (this *ReadAllAction) RunPost(params struct{}) {
	defer this.CreateLogInfo(codes.IPItem_LogReadAllIPItems)

	_, err := this.RPC().IPItemRPC().UpdateIPItemsRead(this.AdminContext(), &pb.UpdateIPItemsReadRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧菜单Badge更新
	helpers.NotifyIPItemsCountChanges()

	this.Success()
}
