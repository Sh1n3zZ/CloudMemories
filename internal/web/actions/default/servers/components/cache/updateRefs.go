// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package cache

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type UpdateRefsAction struct {
	actionutils.ParentAction
}

func (this *UpdateRefsAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdateRefsAction) RunPost(params struct {
	CachePolicyId int64
	RefsJSON      []byte
}) {
	// 修改缓存条件
	if params.CachePolicyId > 0 && len(params.RefsJSON) > 0 {
		_, err := this.RPC().HTTPCachePolicyRPC().UpdateHTTPCachePolicyRefs(this.AdminContext(), &pb.UpdateHTTPCachePolicyRefsRequest{
			HttpCachePolicyId: params.CachePolicyId,
			RefsJSON:          params.RefsJSON,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	this.Success()
}
