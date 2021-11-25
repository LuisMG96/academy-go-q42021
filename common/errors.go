package common

import "net/http"

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func New(err error) errorResponse {
	status, message := GetStatusErrorCode(err)
	return errorResponse{
		Status:  status,
		Message: message,
	}
}

func GetStatusErrorCode(err error) (int, string) {
	switch err.Error() {
	case "500":
		return http.StatusInternalServerError, "Internal Server Error"
	case "400":
		return http.StatusBadRequest, "Bad Request"
	case "5001":
		return http.StatusInternalServerError, "File not Found"
	case "5002":
		return http.StatusInternalServerError, "Malformed File"
	case "5003":
		return http.StatusNotFound, "Character not found"
	case "403":
		return http.StatusForbidden, "Forbidden"
	default:
		return http.StatusInternalServerError, "Internal Server Error"
	}
}
