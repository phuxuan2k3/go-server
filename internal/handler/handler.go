package handler

import (
	hello "darius/pkg/proto/hello"
	suggest "darius/pkg/proto/suggest"
)

type Dependency struct {
}

type handler struct {
	hello.HelloServiceServer
	suggest.SuggestServiceServer
}

// mustEmbedUnimplementedHelloServiceServer implements hello.HelloServiceServer.
func (h *handler) mustEmbedUnimplementedHelloServiceServer() {
	panic("unimplemented")
}

func NewHandlerWithDeps(deps Dependency) *handler {
	return &handler{}
}
