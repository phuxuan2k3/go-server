package cmd

import (
	"context"
	"darius/pkg/proto/hello"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}

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

	grpcPort := viper.GetString("GRPC_PORT")

	err := hello.RegisterHelloServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+grpcPort, opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	gatewayPort := viper.GetString("GATEWAY_PORT")
	log.Println("HTTP Gateway running on port " + gatewayPort)

	server := &http.Server{
		Addr:    ":" + gatewayPort,
		Handler: corsMiddleware(mux),
	}
	server.ListenAndServe()
}
