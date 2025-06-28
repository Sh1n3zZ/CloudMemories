package admins

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	AdminId int64
}) {
	defer this.CreateLogInfo(codes.Admin_LogDeleteAdmin, params.AdminId)

	_, err := this.RPC().AdminRPC().DeleteAdmin(this.AdminContext(), &pb.DeleteAdminRequest{AdminId: params.AdminId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知更改
	err = configloaders.NotifyAdminModuleMappingChange()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
