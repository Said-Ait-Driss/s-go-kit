package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeUppercaseEndpoint(srv StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uppercaseRequest)
		v, err := srv.Uppercase(req.S)

		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}

		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(srv StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(countRequest)
		v := srv.Count(req.S)
		return countResponse{v}, nil
	}
}
