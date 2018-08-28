// Package lorem_grpc : need to encode / decode the request and response.
// For this purpose, create model.go and define encode/decode functions.
package lorem_grpc

import (
	"context"

	"github.com/iqlaas/lorem-grpc/pb"
)

// EncodeGRPCLoremRequest Encode and Decode Lorem Request
func EncodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(LoremRequest)
	return &pb.LoremRequest{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

// DecodeGRPCLoremRequest function
func DecodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoremRequest)
	return LoremRequest{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

// EncodeGRPCLoremResponse function
func EncodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(LoremResponse)
	return &pb.LoremResponse{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}

// DecodeGRPCLoremResponse function
func DecodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.LoremResponse)
	return LoremResponse{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}
