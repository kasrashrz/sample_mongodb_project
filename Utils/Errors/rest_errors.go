package Errors
//TODO: CREATE REST ERRORS AND USE IT ANYWHERE!!

type RestError struct {
	Message string	`json:"message"`
	Status 	int		`json:"code"`
	Error 	string	`json:"error"`
}