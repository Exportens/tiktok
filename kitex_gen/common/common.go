package common

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Err int64

const (
	Err_BadRequest           Err = 10001
	Err_Unauthorized         Err = 10002
	Err_ServerNotFound       Err = 10003
	Err_ServerMethodNotFound Err = 10004
	Err_RequestServerFail    Err = 10005
	Err_ServerHandleFail     Err = 10006
	Err_ResponseUnableParse  Err = 10007
	Err_DuplicateOutOrderNo  Err = 20001
)

func (p Err) String() string {
	switch p {
	case Err_BadRequest:
		return "BadRequest"
	case Err_Unauthorized:
		return "Unauthorized"
	case Err_ServerNotFound:
		return "ServerNotFound"
	case Err_ServerMethodNotFound:
		return "ServerMethodNotFound"
	case Err_RequestServerFail:
		return "RequestServerFail"
	case Err_ServerHandleFail:
		return "ServerHandleFail"
	case Err_ResponseUnableParse:
		return "ResponseUnableParse"
	case Err_DuplicateOutOrderNo:
		return "DuplicateOutOrderNo"
	}
	return "<UNSET>"
}

func ErrFromString(s string) (Err, error) {
	switch s {
	case "BadRequest":
		return Err_BadRequest, nil
	case "Unauthorized":
		return Err_Unauthorized, nil
	case "ServerNotFound":
		return Err_ServerNotFound, nil
	case "ServerMethodNotFound":
		return Err_ServerMethodNotFound, nil
	case "RequestServerFail":
		return Err_RequestServerFail, nil
	case "ServerHandleFail":
		return Err_ServerHandleFail, nil
	case "ResponseUnableParse":
		return Err_ResponseUnableParse, nil
	case "DuplicateOutOrderNo":
		return Err_DuplicateOutOrderNo, nil
	}
	return Err(0), fmt.Errorf("not a valid Err string")
}

func ErrPtr(v Err) *Err { return &v }

func (p *Err) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = Err(result.Int64)
	return
}

func (p *Err) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}
