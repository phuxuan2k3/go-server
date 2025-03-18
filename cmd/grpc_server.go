package cmd

import (
	"darius/internal/handler"
	"darius/internal/llm"
	hello "darius/pkg/proto/hello"
	suggest "darius/pkg/proto/suggest"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func startGRPC() {
	port := viper.GetString("grpc.port")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// fmt.Println("get env config.yaml" + viper.GetString("llm.host"))
	// fmt.Println("get env config.yaml" + viper.GetString("llm.model"))

	llmHost := viper.GetString("llm.host")
	if llmHost == "" || !strings.HasPrefix(llmHost, "http") {
		llmHost = "http://104.199.250.71:2525/api/chat/completions"
	}
	llmModel := viper.GetString("llm.model")
	if llmModel == "" || strings.HasPrefix(llmModel, "$") {
		llmModel = "gpt-4o-mini"
	}
	fmt.Println("llmHost: " + llmHost)
	fmt.Println("llmModel: " + llmModel)
	LlmService := llm.NewLLM(&llm.Config{
		Host:  llmHost,
		Model: llmModel,
	})

	handler := handler.NewHandlerWithDeps(handler.Dependency{
		LlmService: LlmService,
	})

	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, handler)
	suggest.RegisterSuggestServiceServer(grpcServer, handler)

	fmt.Println("gRPC server listening on port " + port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
