package main

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

type MyEvent struct {
	Payload string `json:"payload"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, event MyEvent) (Response, error) {
	switch event.Payload {
	case "panic":
		panic("error")
	case "error":
		return Response{}, errors.New("error")
	case "print":
		println("print caused by event")
	case "stdout":
		os.Stdout.Write([]byte("stdout caused by event"))
	case "stderr":
		os.Stderr.Write([]byte("stderr caused by event"))
	}
	return Response{Message: "Hello World"}, nil
}
