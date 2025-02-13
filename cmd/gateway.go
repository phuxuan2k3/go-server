package cmd

import (
	"github.com/spf13/cobra"

	"context"
	"log"
	"net/http"

	hello "darius/pkg/proto/hello"
	suggest "darius/pkg/proto/suggest"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var gateway = &cobra.Command{
	Use:   "gateway",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.DialContext(
			context.Background(),
			"darius-grpc:50051", // if you run this locally, change this to localhost:50051
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("Failed to dial server: %v", err)
		}

		mux := runtime.NewServeMux()

		err = hello.RegisterHelloServiceHandler(context.Background(), mux, conn)
		if err != nil {
			log.Fatalf("Failed to register gateway: %v", err)
		}
		err = suggest.RegisterSuggestServiceHandler(context.Background(), mux, conn)
		if err != nil {
			log.Fatalf("Failed to register gateway: %v", err)
		}

		gwServer := &http.Server{
			Addr:    ":8080",
			Handler: mux,
		}

		log.Println("Starting HTTP server on port 8080")
		log.Fatal(gwServer.ListenAndServe())
	},
}
