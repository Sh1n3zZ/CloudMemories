package accesskeys

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type UpdateIsOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateIsOnAction) RunPost(params struct {
	AccessKeyId int64
	IsOn        bool
}) {
	defer this.CreateLogInfo(codes.UserAccessKey_LogUpdateUserAccessKeyIsOn, params.AccessKeyId)

	_, err := this.RPC().UserAccessKeyRPC().UpdateUserAccessKeyIsOn(this.AdminContext(), &pb.UpdateUserAccessKeyIsOnRequest{
		UserAccessKeyId: params.AccessKeyId,
		IsOn:            params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
