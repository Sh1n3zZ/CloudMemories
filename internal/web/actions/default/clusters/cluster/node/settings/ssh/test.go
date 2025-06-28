// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ssh

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type TestAction struct {
	actionutils.ParentAction
}

func (this *TestAction) RunPost(params struct {
	GrantId int64
	Host    string
	Port    int32
}) {
	resp, err := this.RPC().NodeGrantRPC().TestNodeGrant(this.AdminContext(), &pb.TestNodeGrantRequest{
		NodeGrantId: params.GrantId,
		Host:        params.Host,
		Port:        params.Port,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["isOk"] = resp.IsOk
	this.Data["error"] = resp.Error
	this.Success()
}
