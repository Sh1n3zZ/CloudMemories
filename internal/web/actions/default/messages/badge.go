package messages

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type BadgeAction struct {
	actionutils.ParentAction
}

func (this *BadgeAction) RunPost(params struct{}) {
	countResp, err := this.RPC().MessageRPC().CountUnreadMessages(this.AdminContext(), &pb.CountUnreadMessagesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["count"] = countResp.Count

	this.Success()
}
