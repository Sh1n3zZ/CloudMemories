package database

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DeleteTableAction struct {
	actionutils.ParentAction
}

func (this *DeleteTableAction) RunPost(params struct {
	Table string
}) {
	defer this.CreateLogInfo(codes.Database_LogDeleteTable, params.Table)

	_, err := this.RPC().DBRPC().DeleteDBTable(this.AdminContext(), &pb.DeleteDBTableRequest{DbTable: params.Table})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
