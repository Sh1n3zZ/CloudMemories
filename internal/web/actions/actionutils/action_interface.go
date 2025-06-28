package actionutils

import (
	"context"

	"github.com/Sh1n3zZ/CloudMemories/internal/rpc"
	"github.com/iwind/TeaGo/maps"
)

type ActionInterface interface {
	RPC() *rpc.RPCClient

	AdminContext() context.Context

	ViewData() maps.Map
}
