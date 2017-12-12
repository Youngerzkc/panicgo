package errors

import (
	"fmt"
)

type PanicError struct {
	ErrMsg string	`json:"errmsg"`
	ErrNo    int	`json:"errno"`
}

func Wrap(err error, errno int, errmsg string) *PanicError {
	//fmt.Errorf("%v", err)
	return &PanicError{
		ErrNo:    errno,
		ErrMsg: errmsg,
	}
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.ErrNo, e.ErrMsg)
}
