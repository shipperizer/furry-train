package health

import (
	"google.golang.org/grpc"
)

// RPCInterface interface for RPC.
type RPCInterface interface {
	Register(*grpc.Server)
}
