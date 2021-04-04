package Errors

import "net/http"

//TODO: CREATE REST ERRORS AND USE IT ANYWHERE!!

type RestError struct {
	Message string	`json:"message"`
	Status 	int		`json:"code"`
	Error 	string	`json:"error"`
}

func BadRequest (message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad_Request",
	}
}

func NotFoundError (message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not_Found",
	}
}

func ServerError (message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal_server_Error",
	}
}