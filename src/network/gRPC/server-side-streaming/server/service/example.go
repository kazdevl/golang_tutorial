package service

import (
	"app/network/gRPC/server-side-streaming/pb"
	"context"
	"fmt"
	"log"
)

var messageQueue = make(chan string)

type Example struct {
	pb.UnimplementedExampleServiceServer
}

func NewExample() *Example {
	return &Example{}
}

func (e *Example) GetCount(_ *pb.GetCountRequest, stream pb.ExampleService_GetCountServer) error {
	count := 0
	for msg := range messageQueue {
		if len(msg) == 0 {
			return nil
		}

		count += 1
		if err := stream.Send(&pb.GetCountResponse{Message: fmt.Sprintf("GetMessage[%d回目]: %s", count, msg)}); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (e *Example) Push(ctx context.Context, in *pb.PushRequest) (*pb.PushResponse, error) {
	log.Println("Call Push")
	messageQueue <- fmt.Sprintf("pushd: %s", in.Message)
	return &pb.PushResponse{}, nil
}
