package errors

/**
* Struct to manage Rest Errors
 */
type RestError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
