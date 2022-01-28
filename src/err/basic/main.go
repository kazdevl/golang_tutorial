package main

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	fmt.Printf("%+v\n", ErrRoot())
}

func ErrRoot() error {
	return xerrors.Errorf("err root:%w", ErrDepthOne())
}

func ErrDepthOne() error {
	return fmt.Errorf("error depth1:%w", ErrDepthTwo())
}

func ErrDepthTwo() error {
	return errors.New("error depth2")
}
