package check

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Blueprint struct for db service and auth provider.
type RPC struct {
	UnimplementedCheckServiceServer
}

// Check applies the logic to determine which variant to be applied
// to the given request.
func (rpc *RPC) Check(context.Context, *emptypb.Empty) (*CheckResponse, error) {
	return &CheckResponse{Status: http.StatusOK, Message: "purple bro"}, nil
}

// Register wires up the RPC to the server
func (rpc *RPC) Register(server *grpc.Server) {
	RegisterCheckServiceServer(server, rpc)
}

// NewRPC returns a new initialized RPC object.
func NewRPC() RPCInterface {
	return new(RPC)
}
