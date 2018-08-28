package client

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	lorem_grpc "github.com/iqlaas/lorem-grpc"
	"github.com/iqlaas/lorem-grpc/pb"
	"google.golang.org/grpc"
)

// New return new lorem_grpc service
func New(conn *grpc.ClientConn) lorem_grpc.Service {
	var loremEndpoint = grpctransport.NewClient(
		conn, "Lorem", "Lorem",
		lorem_grpc.EncodeGRPCLoremRequest,
		lorem_grpc.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()
	return lorem_grpc.Endpoints{
		LoremEndpoint: loremEndpoint,
	}
}
