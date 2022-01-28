package main

import (
	"fmt"

	"github.com/friendsofgo/errors"
)

type ErrClient struct {
	Err error
}

func (e ErrClient) CallErr(message string) error {
	return errors.WithStack(e.Err)
}

type Sample struct {
	EClient ErrClient
}

func (s Sample) GetError() error {
	return errors.WithStack(s.EClient.CallErr("sample"))
}

func main() {
	e := ErrClient{Err: errors.New("error1")}
	s := Sample{EClient: e}
	fmt.Printf("%+v\n", s.GetError())
}
