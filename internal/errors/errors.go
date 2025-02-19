package errors

type ErrorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type HttpError struct {
	StatusCode int
	Code       string
	Message    string
	Details    interface{}
}

type ServiceError struct {
	Code    string
	Message string
	Details interface{}
	Cause   error
}

func (e *HttpError) Error() string {
	return e.Message
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewHttpError(statusCode int, code, message string, details interface{}) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Details:    details,
	}
}

func NewServiceError(code, message string, details interface{}, cause error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
		Details: details,
		Cause:   cause,
	}
}
