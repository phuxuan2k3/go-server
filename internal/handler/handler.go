package handler

import (
	llm "darius/internal/llm"
	hello "darius/pkg/proto/hello"
	suggest "darius/pkg/proto/suggest"
)

type Dependency struct {
	LlmService llm.LLM
}

type handler struct {
	hello.HelloServiceServer
	suggest.SuggestServiceServer

	llmService llm.LLM
}

func NewHandlerWithDeps(deps Dependency) *handler {
	return &handler{
		llmService: deps.LlmService,
	}
}
