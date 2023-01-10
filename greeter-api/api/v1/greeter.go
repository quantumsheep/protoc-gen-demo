package v1

import (
	"context"
	"fmt"

	greeterv1 "github.com/scaleway/protobuf/go/demo/greeter/v1"
)

type GreeterServer struct {
	greeterv1.UnimplementedGreeterServiceServer
}

func (h *GreeterServer) SayHello(ctx context.Context, request *greeterv1.SayHelloRequest) (*greeterv1.SayHelloResponse, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, err
	}

	fmt.Println(request.GetName(), "says hello")

	return &greeterv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s!", request.GetName()),
	}, nil
}

func (h *GreeterServer) SayGoodbye(ctx context.Context, request *greeterv1.SayGoodbyeRequest) (*greeterv1.SayGoodbyeResponse, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, err
	}

	fmt.Println(request.GetName(), "says goodbye")

	return &greeterv1.SayGoodbyeResponse{
		Message: fmt.Sprintf("Goodbye %s!", request.GetName()),
	}, nil
}
