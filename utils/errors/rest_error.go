package errors

import "net/http"

/**
* Struct to manage Rest Errors
 */
type RestError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type NewBadRequestError struct {
	RestError RestError
}

func (newBadRequestError *NewBadRequestError) NewBadRequest(message string) *RestError {
	newBadRequestError.RestError.Error = "bad_request"
	newBadRequestError.RestError.Message = message
	newBadRequestError.RestError.Status = http.StatusBadRequest
	return &newBadRequestError.RestError
}
