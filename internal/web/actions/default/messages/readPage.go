package messages

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type ReadPageAction struct {
	actionutils.ParentAction
}

func (this *ReadPageAction) RunPost(params struct {
	MessageIds []int64
}) {
	// 创建日志
	defer this.CreateLogInfo(codes.Message_LogReadMessages)

	_, err := this.RPC().MessageRPC().UpdateMessagesRead(this.AdminContext(), &pb.UpdateMessagesReadRequest{
		MessageIds: params.MessageIds,
		IsRead:     true,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
