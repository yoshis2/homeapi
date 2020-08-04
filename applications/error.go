package applications

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

//UError Usecase Error
type UsecaseError struct {
	Code int
	Msg  string
}

const (
	BadRequest        = iota + 1 // BadRequest リクエスト不正
	Unauthorized                 // Unauthorized 認証不正
	NotFound                     // NotFound 存在しない
	Conflict                     // Conflict 競合
	Internalexception            // Internalexception その他内部エラー
)

//Error error interfaceを実装
func (err *UsecaseError) Error() string {
	return fmt.Sprintf("ERROR: %s", err.Msg)
}

//GetUError Usecase Error
func GetUError(code int, msg string) *UsecaseError {
	uerr := &UsecaseError{
		Code: code,
		Msg:  msg,
	}

	return uerr
}

//GetUErrorByError Get Usecase Error By Error
func GetUErrorByError(err error) *UsecaseError {
	if err != nil {
		var code int
		if gorm.IsRecordNotFoundError(err) {
			code = NotFound
		} else if isDuplicatedUError(err.Error()) {
			code = Conflict
		} else {
			code = Internalexception
		}

		return GetUError(code, err.Error())
	}
	return nil
}

//IsRecordNotFoundUError Error is NotFound
func (err *UsecaseError) IsRecordNotFoundUError() bool {
	return err.Code == NotFound
}

func isDuplicatedUError(msg string) bool {
	return strings.Index(msg, "Duplicate") != -1
}
