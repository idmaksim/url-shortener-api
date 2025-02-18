package httpErrors

type HTTPError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(message string, status int) *HTTPError {
	return &HTTPError{Message: message, Status: status}
}
