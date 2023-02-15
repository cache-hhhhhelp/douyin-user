package xerr

import (
	"fmt"
)

type CodeError struct {
	StatusCode uint32
	StatusMsg  string
}

func (e *CodeError) GetErrCode() uint32 {
	return e.StatusCode
}

func (e *CodeError) GetErrMsg() string {
	return e.StatusMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("StatusCode:%d，StatusMsg:%s", e.StatusCode, e.StatusMsg)
}

func NewErrCodeMsg(StatusCode uint32, StatusMsg string) *CodeError {
	return &CodeError{
		StatusCode: StatusCode,
		StatusMsg:  StatusMsg,
	}
}
