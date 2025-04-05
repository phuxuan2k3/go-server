package cmd

import (
	"darius/internal/handler"

	hello "darius/pkg/proto/hello"
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func startGRPC() {
	//server gateway
	port := viper.GetString("GRPC_PORT")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//get llm host and model
	// llmHost := viper.GetString("llm.host")
	// if llmHost == "" || !strings.HasPrefix(llmHost, "http") {
	// 	llmHost = "http://104.199.250.71:2525/api/chat/completions"
	// }
	// llmModel := viper.GetString("llm.model")
	// if llmModel == "" || strings.HasPrefix(llmModel, "$") {
	// 	llmModel = "gpt-4o-mini"
	// }

	// llmGRPCAddress := viper.GetString("llm_grpc.host")
	// if llmGRPCAddress == "" || strings.HasPrefix(llmGRPCAddress, "$") {
	// 	llmGRPCAddress = "104.199.250.71"
	// }
	// llmGRPCPort := viper.GetString("llm_grpc.port")
	// if llmGRPCPort == "" || strings.HasPrefix(llmGRPCPort, "$") {
	// 	llmGRPCPort = "2524"
	// }
	// llmGRPCModel := viper.GetString("llm.model")
	// if llmGRPCModel == "" || strings.HasPrefix(llmGRPCModel, "$") {
	// 	llmGRPCModel = "gpt-4o-mini"
	// }
	// addr := flag.String("addr", llmGRPCAddress+":"+llmGRPCPort, "the address to connect to")
	// flag.Parse()
	// conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()

	// arceusClient := suggest.NewArceusClient(conn)
	// llmGRPCService := llm_grpc.NewService(arceusClient, llmGRPCModel)

	// r, err := c.GenerateText(context.Background(),
	// 	&suggest.GenerateTextRequest{
	// 		Model:   llmModel,
	// 		Content: "Hello, how are you?"})

	// LlmService := llm.NewLLM(&llm.Config{
	// 	Host:  llmHost,
	// 	Model: llmModel,
	// })
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.Content)

	handler := handler.NewHandlerWithDeps(handler.Dependency{
		// LlmService: LlmService,
		// LLMGRPC: llmGRPCService,
		LLMGRPC: nil,
	})

	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, handler)
	// suggest.RegisterSuggestServiceServer(grpcServer, handler)

	fmt.Print("This is the log message that prove that the server is redeployed by owner Xuan\n")

	fmt.Println("gRPC server listening on port " + port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
