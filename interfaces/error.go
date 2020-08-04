package interfaces

import (
	"net/http"

	"homeapi/applications"
)

//ErrorResponseObject error時に返却するオブジェクト
type ErrorResponseObject struct {
	Message string `json:"message"`
}

//GetErrorResponse ErrorCodeとErrorResponseObjectを返却する
func GetErrorResponse(err *applications.UsecaseError) (int, ErrorResponseObject) {
	return getErrorHTTPStatus(err.Code), ErrorResponseObject{
		Message: err.Msg,
	}
}

func getErrorHTTPStatus(errCode int) int {
	switch errCode {
	case applications.BadRequest:
		return http.StatusBadRequest
	case applications.Unauthorized:
		return http.StatusUnauthorized
	case applications.NotFound:
		return http.StatusNotFound
	case applications.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
