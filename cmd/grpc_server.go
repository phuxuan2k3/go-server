package cmd

import (
	"darius/internal/handler"
	"darius/internal/llm"
	hello "darius/pkg/proto/hello"
	suggest "darius/pkg/proto/suggest"
	"fmt"
	"log"
	"net"

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

	LlmService := llm.NewLLM(&llm.Config{
		Host: viper.GetString("llm.host"),
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
