package acme

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteTaskAction struct {
	actionutils.ParentAction
}

func (this *DeleteTaskAction) RunPost(params struct {
	TaskId int64
}) {
	defer this.CreateLogInfo(codes.ACMETask_LogDeleteACMETask, params.TaskId)

	_, err := this.RPC().ACMETaskRPC().DeleteACMETask(this.AdminContext(), &pb.DeleteACMETaskRequest{AcmeTaskId: params.TaskId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
