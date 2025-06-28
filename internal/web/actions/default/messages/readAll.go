package messages

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type ReadAllAction struct {
	actionutils.ParentAction
}

func (this *ReadAllAction) RunPost(params struct{}) {
	// 创建日志
	defer this.CreateLogInfo(codes.Message_LogReadAll)

	_, err := this.RPC().MessageRPC().UpdateAllMessagesRead(this.AdminContext(), &pb.UpdateAllMessagesReadRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
