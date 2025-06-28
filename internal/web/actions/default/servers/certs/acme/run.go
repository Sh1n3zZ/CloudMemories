package acme

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type RunAction struct {
	actionutils.ParentAction
}

func (this *RunAction) RunPost(params struct {
	TaskId int64
}) {
	defer this.CreateLogInfo(codes.ACMETask_LogRunACMETask, params.TaskId)

	runResp, err := this.RPC().ACMETaskRPC().RunACMETask(this.AdminContext(), &pb.RunACMETaskRequest{AcmeTaskId: params.TaskId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if runResp.IsOk {
		this.Data["certId"] = runResp.SslCertId
		this.Success()
	} else {
		this.Fail(runResp.Error)
	}
}
