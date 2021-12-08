package common

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//New - Constructor of errorResponse Struct who receives an error, parse it and returns it as an errorResponse
func NewResponse(status int, message string) response {
	return response{
		Status:  status,
		Message: message,
	}

}
