package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kazdevl/golang_tutorial/network/gRPC/server-side-streaming/pb"
	"github.com/kazdevl/golang_tutorial/network/gRPC/server-side-streaming/server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8080
	listerPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	pb.RegisterExampleServiceServer(srv, service.NewExample())
	reflection.Register(srv)
	srv.Serve(listerPort)
}
