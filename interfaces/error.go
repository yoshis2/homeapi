package interfaces

import (
	"log"
	"net/http"
	"strings"
)

//ErrorResponseObject error時に返却するオブジェクト
type ErrorResponseObject struct {
	Message string `json:"message"`
}

// ErrorResponse はエラーレスポンス
func ErrorResponse(err error) (int, ErrorResponseObject) {
	var code int
	if err != nil {
		if isBadRequestError(err.Error()) {
			code = http.StatusBadRequest
		} else if isDuplicatedUError(err.Error()) {
			code = http.StatusConflict
		} else {
			code = http.StatusInternalServerError
		}
	}
	return code, ErrorResponseObject{
		Message: err.Error(),
	}
}

func isDuplicatedUError(msg string) bool {
	return strings.Contains(msg, "Duplicate")
}

func isBadRequestError(msg string) bool {
	var messageBool bool

	log.Printf("message : %v", msg)
	if strings.Contains(msg, "a foreign key constraint fails") {
		messageBool = true
	} else if strings.Contains(msg, "BadRequest") {
		messageBool = true
	}
	return messageBool
}
