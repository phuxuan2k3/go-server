package cmd

import (
	"context"
	suggest "darius/pkg/proto/suggest"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func startGateway() {
	// conn, err := grpc.DialContext(
	// 	context.Background(),
	// 	"darius-grpc:50051", // if you run this locally, change this to localhost:50051
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to dial server: %v", err)
	// }

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// err := hello.RegisterHelloServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	// if err != nil {
	// 	log.Fatalf("Failed to register gateway: %v", err)
	// }

	grpcPort := viper.GetString("grpc.port")

	err := suggest.RegisterSuggestServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+grpcPort, opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	gatewayPort := viper.GetString("gateway.port")
	log.Println("HTTP Gateway running on port " + gatewayPort)
	http.ListenAndServe(":"+gatewayPort, mux)
}
