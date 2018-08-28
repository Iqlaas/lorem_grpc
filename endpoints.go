package lorem_grpc

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// LoremRequest struct
type LoremRequest struct {
	RequestType string
	Min         int32
	Max         int32
}

// LoremResponse struct
type LoremResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

// Endpoints struct
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

// MakeLoremEndpoint endpoint
func MakeLoremEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)

		var (
			min, max int
		)

		min = int(req.Min)
		max = int(req.Max)
		txt, err := svc.Lorem(ctx, req.RequestType, min, max)

		if err != nil {
			return nil, err
		}
		return LoremResponse{Message: txt}, nil
	}
}

// Lorem .. : Wrapping endpoints as a Service implementation
// Will be used in grpc client
func (e Endpoints) Lorem(ctx context.Context, requestType string, min, max int) (string, error) {
	req := LoremRequest{
		RequestType: requestType,
		Min:         int32(min),
		Max:         int32(max),
	}
	resp, err := e.LoremEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	loremResp := resp.(LoremResponse)
	if loremResp.Err != "" {
		return "", errors.New(loremResp.Err)
	}
	return loremResp.Message, nil
}
