package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yuyuancha/go-grpc-practice/hello"
	"google.golang.org/grpc"
)

func main() {
	var gRPCServer = grpc.NewServer()
	hello.RegisterHelloServiceServer(gRPCServer, new(hello.HelloServer))

	var listen, err = net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started...")
	log.Fatal(gRPCServer.Serve(listen))
}
