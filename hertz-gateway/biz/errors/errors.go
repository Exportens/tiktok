package errors

import (
	"api_gateway/kitex_gen/common"
)

type Err struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// New Error, the error_code must be defined in IDL.
func New(errCode common.Err) Err {
	return Err{
		ErrCode: int64(errCode),
		ErrMsg:  errCode.String(),
	}
}

func (e Err) Error() string {
	return e.ErrMsg
}
