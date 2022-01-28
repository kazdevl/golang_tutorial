package main

import (
	"fmt"

	"github.com/friendsofgo/errors"
	// "github.com/pkg/errors"

	"golang.org/x/xerrors"
)

func main() {
	fmt.Printf("%+v\n", ErrRoot())
}

func ErrRoot() error {
	return xerrors.Errorf("err root: %w", ErrDepthOne())
}

func ErrDepthOne() error {
	return xerrors.Errorf("error depth1: %w", ErrDepthTwo())
}

func ErrDepthTwo() error {
	return errors.Wrap(fmt.Errorf("sample: %s", "data"), "deepest error")
}
