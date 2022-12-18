package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kazdevl/golang_tutorial/network/gRPC/mock-server/pb"
	"google.golang.org/grpc"
)

func main() {
	address := "g-15v8pmnkpknm5mklwom5lgd8zom6om.srv.pstmn.io"
	conn, err := grpc.Dial(address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewSampleServiceClient(conn)
	req := &pb.HelloRequest{
		TargetUserId: 1,
		Message:      "Hello",
	}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Message)
}
