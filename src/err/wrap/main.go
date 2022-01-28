package main

import (
	"fmt"

	"github.com/friendsofgo/errors"
)

var (
	ErrorForClient = errors.New("error1")
)

type ErrClient struct {
	Err error
}

func (e ErrClient) CallErr(message string) error {
	return errors.Wrap(e.Err, "client error")
}

type Sample struct {
	EClient ErrClient
}

func (s Sample) GetError() error {
	return errors.Wrap(s.EClient.CallErr("sample"), "Sample error")
}

func main() {
	e := ErrClient{Err: ErrorForClient}
	s := Sample{EClient: e}
	if err := s.GetError(); err != nil {
		if err := errors.Cause(err); errors.Is(ErrorForClient, err) {
			fmt.Println("target error")
			fmt.Printf("%+v\n", err)
			fmt.Printf("%+v\n", errors.New("sample"))
			fmt.Println("****************")
		}
		fmt.Printf("%+v\n", err)
	}
}
