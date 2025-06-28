// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package webp

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/dao"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/groups/group/servergrouputils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("webp")
}

func (this *IndexAction) RunGet(params struct {
	GroupId int64
}) {
	_, err := servergrouputils.InitGroup(this.Parent(), params.GroupId, "webp")
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
	this.Data["webpConfig"] = webConfig.WebP

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId    int64
	WebpJSON []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	var webpConfig = &serverconfigs.WebPImageConfig{}
	err := json.Unmarshal(params.WebpJSON, webpConfig)
	if err != nil {
		this.Fail("参数校验失败：" + err.Error())
	}

	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebWebP(this.AdminContext(), &pb.UpdateHTTPWebWebPRequest{
		HttpWebId: params.WebId,
		WebpJSON:  params.WebpJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
