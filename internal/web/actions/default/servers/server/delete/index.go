package deletes

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "delete", "")
	this.SecondMenu("index")
}

func (this *IndexAction) RunGet(params struct{}) {
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId int64
	Must     *actions.Must
}) {
	// 记录日志
	defer this.CreateLogInfo(codes.Server_LogDeleteServer, params.ServerId)

	// 执行删除
	_, err := this.RPC().ServerRPC().DeleteServer(this.AdminContext(), &pb.DeleteServerRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
