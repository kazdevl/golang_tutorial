package perror

import (
	"fmt"

	"github.com/kazdevl/golang_tutorial/err/handling/constant"

	"github.com/friendsofgo/errors"
	"google.golang.org/genproto/googleapis/rpc/code"
)

type ErrorCase struct {
	GRPCErrorCode    code.Code
	ProductErrorCode constant.ErrorCode
	OperationCode    constant.OperationCode
}

var (
	ErrorCase_Unknown = ErrorCase{
		GRPCErrorCode:    code.Code_UNKNOWN,
		ProductErrorCode: constant.ErrorCode_Unknown,
		OperationCode:    constant.OperationCode_Report,
	}
	ErrorCase_FailedDB = ErrorCase{
		GRPCErrorCode:    code.Code_INTERNAL,
		ProductErrorCode: constant.ErrorCode_FailedDB,
		OperationCode:    constant.OperationCode_Report,
	}
	ErrorCase_InvalidToken = ErrorCase{
		GRPCErrorCode:    code.Code_INVALID_ARGUMENT,
		ProductErrorCode: constant.ErrorCode_InvalidToken,
		OperationCode:    constant.OperationCode_Relogin,
	}
)

type ProductError struct {
	ErrorCase ErrorCase
	err       error
}

func (p ProductError) Error() string {
	return fmt.Sprintf("error case: %+v, err: %+v\n", p.ErrorCase, p.err)
}

func (p ProductError) UnWrap() error {
	return p.err
}

func Error(errorCase ErrorCase, message string) error {
	return newError(errorCase, message)
}

func Errorf(errorCase ErrorCase, format string, args ...interface{}) error {
	return newError(errorCase, fmt.Sprintf(format, args...))
}

func newError(errorCase ErrorCase, message string) ProductError {
	return ProductError{
		ErrorCase: errorCase,
		err:       errors.New(message),
	}
}

func Wrap(err error, errorCase ErrorCase) error {
	return newWrap(err, errorCase, "")
}

func Wrapf(err error, errorCase ErrorCase, format string, args ...interface{}) error {
	return newWrap(err, errorCase, fmt.Sprintf(format, args...))
}

func newWrap(err error, errorCase ErrorCase, message string) ProductError {
	return ProductError{
		ErrorCase: errorCase,
		err:       errors.Wrap(err, message),
	}
}

func Stack(err error) error {
	return newStack(err, "")
}

func Stackf(err error, message string) error {
	return newStack(err, message)
}

func newStack(err error, message string) error {
	perr, ok := err.(ProductError)
	if ok {
		perr.err = errors.Wrap(perr.err, message)
		return perr
	}
	return newWrap(err, ErrorCase_Unknown, "unknown")
}
