package tasks

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteBatchAction struct {
	actionutils.ParentAction
}

func (this *DeleteBatchAction) RunPost(params struct {
	TaskIds []int64
}) {
	defer this.CreateLogInfo(codes.NodeTask_LogDeleteNodeTasksBatch)

	_, err := this.RPC().NodeTaskRPC().DeleteNodeTasks(this.AdminContext(), &pb.DeleteNodeTasksRequest{NodeTaskIds: params.TaskIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
