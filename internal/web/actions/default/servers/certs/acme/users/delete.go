package users

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	UserId int64
}) {
	defer this.CreateLogInfo(codes.ACMEUser_LogDeleteACMEUser, params.UserId)

	countResp, err := this.RPC().ACMETaskRPC().CountAllEnabledACMETasksWithACMEUserId(this.AdminContext(), &pb.CountAllEnabledACMETasksWithACMEUserIdRequest{AcmeUserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if countResp.Count > 0 {
		this.Fail("有任务正在和这个用户关联，所以不能删除")
	}

	_, err = this.RPC().ACMEUserRPC().DeleteACMEUser(this.AdminContext(), &pb.DeleteACMEUserRequest{AcmeUserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
