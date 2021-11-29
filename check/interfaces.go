package check

import (
	"google.golang.org/grpc"
)

// RPCInterface interface for RPC.
type RPCInterface interface {
	Register(*grpc.Server)
}
