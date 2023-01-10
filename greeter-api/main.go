package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/scaleway/greeter-api/api/v1"
	greeterv1 "github.com/scaleway/protobuf/go/demo/greeter/v1"
	"github.com/scaleway/scaleway-sdk-go/api/registry/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	Server()
}

func Server() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	greeterv1.RegisterGreeterServiceServer(grpcServer, &v1.GreeterServer{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Println("Starting gRPC server on http://0.0.0.0:8081")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	gmux := runtime.NewServeMux()
	err = greeterv1.RegisterGreeterServiceHandlerFromEndpoint(ctx, gmux, "localhost:8081", []grpc.DialOption{
		// grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	log.Println("Starting HTTP/1.1 server on http://0.0.0.0:8080")
	if err := http.ListenAndServe(":8080", gmux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Bird struct{}
type ListBirdsRequest struct {
	Page       int
	MaxResults int
}

func Client() {
	conn, err := grpc.Dial(
		"localhost:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := greeterv1.NewGreeterServiceClient(conn)

	request := &greeterv1.SayHelloRequest{
		Name: "John",
	}

	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to say hello: %v", err)
	}

	fmt.Println(response.GetMessage())
}

func Demo() {
	// Get Scaleway Config
	config, err := scw.LoadConfig()
	if err != nil {
		log.Panicf("failed to load config: %s", err)
	}

	// Use active profile
	profile, err := config.GetActiveProfile()
	if err != nil {
		log.Panicf("failed to get active profile: %s", err)
	}

	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithProfile(profile),
		scw.WithEnv(), // env variable may overwrite profile values
	)
	if err != nil {
		log.Panicf("failed to create client: %s", err)
	}

	api := registry.NewAPI(client)

	namespace, err := api.CreateNamespace(
		&registry.CreateNamespaceRequest{},
	)

	if err != nil {
		log.Panicf("failed to create namespace: %s", err)
	}

	namespace, err = api.WaitForNamespace(&registry.WaitForNamespaceRequest{
		NamespaceID: namespace.ID,
	})
	if err != nil {
		log.Panicf("failed to wait for namespace: %s", err)
	}

	fmt.Println(namespace)
}
