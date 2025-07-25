package group

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	GroupId int64
}) {
	// 创建日志
	defer this.CreateLogInfo(codes.ServerGroup_LogDeleteServerGroup, params.GroupId)

	// 检查是否正在使用
	countResp, err := this.RPC().ServerRPC().CountAllEnabledServersWithServerGroupId(this.AdminContext(), &pb.CountAllEnabledServersWithServerGroupIdRequest{ServerGroupId: params.GroupId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if countResp.Count > 0 {
		this.Fail("此分组正在被使用不能删除，请修改相关服务后再删除")
	}

	_, err = this.RPC().ServerGroupRPC().DeleteServerGroup(this.AdminContext(), &pb.DeleteServerGroupRequest{ServerGroupId: params.GroupId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
