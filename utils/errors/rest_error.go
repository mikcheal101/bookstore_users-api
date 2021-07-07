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

// Method to throw bad request error
func (newBadRequestError *NewBadRequestError) NewBadRequest(message string) *RestError {
	newBadRequestError.RestError.Error = "bad_request"
	newBadRequestError.RestError.Message = message
	newBadRequestError.RestError.Status = http.StatusBadRequest
	return &newBadRequestError.RestError
}

// Method to throw not found error
func (newBadRequestError *NewBadRequestError) NotFoundRequest(message string) *RestError {
	newBadRequestError.RestError.Error = "not_found"
	newBadRequestError.RestError.Message = message
	newBadRequestError.RestError.Status = http.StatusNotFound
	return &newBadRequestError.RestError
}
