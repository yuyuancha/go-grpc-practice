package hello

import (
	"context"
	"fmt"
	"log"
)

type HelloServer struct{}

func (s *HelloServer) mustEmbedUnimplementedHelloServiceServer() {
	panic("implement me")
}

func (s *HelloServer) SayHello(ctx context.Context, request *SayHelloRequest) (*SayHelloResponse, error) {
	log.Printf("收到名字: %s", request.Name)

	return &SayHelloResponse{Message: fmt.Sprintf("Hello, %s. Nice to meet you!", request.Name)}, nil
}
