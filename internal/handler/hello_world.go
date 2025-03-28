package handler

import (
	"context"
	hello "darius/pkg/proto/hello"
)

func (h *handler) HelloWorld(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message: "Hello World - Xuân"}, nil
}
