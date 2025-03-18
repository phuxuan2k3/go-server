package handler

import (
	llm "darius/internal/llm"
	llm_grpc "darius/internal/llm-grpc"
	hello "darius/pkg/proto/hello"
	suggest "darius/pkg/proto/suggest"
)

type Dependency struct {
	LlmService llm.LLM
	LLMGRPC    llm_grpc.Service
}

type handler struct {
	hello.HelloServiceServer
	suggest.SuggestServiceServer

	llmService     llm.LLM
	llmGRPCService llm_grpc.Service
}

func NewHandlerWithDeps(deps Dependency) *handler {
	return &handler{
		llmService:     deps.LlmService,
		llmGRPCService: deps.LLMGRPC,
	}
}
