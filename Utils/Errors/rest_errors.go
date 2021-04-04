package Errors

import "net/http"

//TODO: CREATE REST ERRORS AND USE IT ANYWHERE!!

type RestError struct {
	Message string	`json:"message"`
	Status 	int		`json:"code"`
	Error 	string	`json:"error"`
}

func BadRequest(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad_Request",
	}
}