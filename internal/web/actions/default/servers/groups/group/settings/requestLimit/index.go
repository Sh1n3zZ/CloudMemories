// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package requestlimit

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/dao"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/groups/group/servergrouputils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("requestLimit")
}

func (this *IndexAction) RunGet(params struct {
	GroupId int64
}) {
	_, err := servergrouputils.InitGroup(this.Parent(), params.GroupId, "requestLimit")
	if err != nil {
		this.ErrorPage(err)
		return
	}

	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerGroupId(this.AdminContext(), params.GroupId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["requestLimitConfig"] = webConfig.RequestLimit

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId            int64
	RequestLimitJSON []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo(codes.ServerRequestLimit_LogUpdateRequestLimitSettings, params.WebId)

	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWebRequestLimit(this.AdminContext(), &pb.UpdateHTTPWebRequestLimitRequest{
		HttpWebId:        params.WebId,
		RequestLimitJSON: params.RequestLimitJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
