package db

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type TruncateTableAction struct {
	actionutils.ParentAction
}

func (this *TruncateTableAction) RunPost(params struct {
	NodeId int64
	Table  string
}) {
	defer this.CreateLogInfo(codes.DBNode_LogTruncateTable, params.NodeId, params.Table)

	_, err := this.RPC().DBNodeRPC().TruncateDBNodeTable(this.AdminContext(), &pb.TruncateDBNodeTableRequest{
		DbNodeId:    params.NodeId,
		DbNodeTable: params.Table,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
