package handler

import (
	"context"
	hello "darius/pkg/proto/hello"
)

func (h *handler) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message: "Hello, " + req.Name}, nil
}
