package errno

import (
	"errors"
)

type ErrNo struct {
	ErrCode  int32  `json:"code"`
	HttpCode int    `json:"http_code"`
	ErrMsg   string `json:"message"`
}

func NewErrNo(code int32, httpCode int, msg string) ErrNo {
	return ErrNo{code, httpCode, msg}
}

func (e ErrNo) Error() string {
	return e.ErrMsg
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func (e ErrNo) WithError(err error) ErrNo {
	e.ErrMsg = e.ErrMsg + ": " + err.Error()
	return e
}

func ConvertErr(err error) ErrNo {
	if err == nil {
		return Success
	}
	if errno, ok := errors.AsType[ErrNo](err); ok {
		return errno
	}

	serviceErr := InternalServerErr
	return serviceErr
}
