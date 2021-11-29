package health

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/grpc"
)

// Blueprint struct for db service and auth provider.
type RPC struct {
	grpc_health_v1.UnimplementedHealthServer
}

// Check applies the logic to determine which variant to be applied
// to the given request.
func (rpc *RPC) Check(ctx context.Context, r *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

// Register wires up the RPC to the server
func (rpc *RPC) Register(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, rpc)
}

// NewRPC returns a new initialized RPC object.
func NewRPC() RPCInterface {
	return new(RPC)
}
