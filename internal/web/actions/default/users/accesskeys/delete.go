package accesskeys

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	AccessKeyId int64
}) {
	defer this.CreateLogInfo(codes.UserAccessKey_LogDeleteUserAccessKey, params.AccessKeyId)

	_, err := this.RPC().UserAccessKeyRPC().DeleteUserAccessKey(this.AdminContext(), &pb.DeleteUserAccessKeyRequest{UserAccessKeyId: params.AccessKeyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
