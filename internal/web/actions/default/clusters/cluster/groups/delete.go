package groups

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
	// 检查是否正在使用
	countResp, err := this.RPC().NodeRPC().CountAllEnabledNodesWithNodeGroupId(this.AdminContext(), &pb.CountAllEnabledNodesWithNodeGroupIdRequest{NodeGroupId: params.GroupId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if countResp.Count > 0 {
		this.Fail("此分组正在被使用不能删除，请修改节点后再删除")
	}

	_, err = this.RPC().NodeGroupRPC().DeleteNodeGroup(this.AdminContext(), &pb.DeleteNodeGroupRequest{NodeGroupId: params.GroupId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLogInfo(codes.NodeGroup_LogDeleteNodeGroup, params.GroupId)

	this.Success()
}
