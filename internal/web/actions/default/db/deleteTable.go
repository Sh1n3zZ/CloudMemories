package db

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteTableAction struct {
	actionutils.ParentAction
}

func (this *DeleteTableAction) RunPost(params struct {
	NodeId int64
	Table  string
}) {
	defer this.CreateLogInfo(codes.DBNode_LogDeleteTable, params.NodeId, params.Table)

	_, err := this.RPC().DBNodeRPC().DeleteDBNodeTable(this.AdminContext(), &pb.DeleteDBNodeTableRequest{
		DbNodeId:    params.NodeId,
		DbNodeTable: params.Table,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
