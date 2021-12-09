package common

import "net/http"

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//New - Constructor of errorResponse Struct who receives an error, parse it and returns it as an errorResponse
func NewError(err error) errorResponse {
	status, message := getStatusErrorCode(err)
	return errorResponse{
		Status:  status,
		Message: message,
	}

}

func getStatusErrorCode(err error) (int, string) {
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
	case "5004":
		return http.StatusBadGateway, "Character not found"
	case "403":
		return http.StatusForbidden, "Forbidden"
	case "401":
		return http.StatusUnauthorized, "Forbidden"
	default:
		return http.StatusInternalServerError, "Internal Server Error"
	}
}
