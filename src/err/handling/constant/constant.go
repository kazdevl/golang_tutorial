package constant

type ErrorCode int

const (
	ErrorCode_Unknown      ErrorCode = 100
	ErrorCode_FailedDB     ErrorCode = 101
	ErrorCode_InvalidToken ErrorCode = 102
)

type OperationCode int

const (
	OperationCode_Report  OperationCode = 200
	OperationCode_Relogin OperationCode = 201
)
