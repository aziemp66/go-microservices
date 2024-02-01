package http_error

func NewError(raw error, code int, message string) *Error {
	return &Error{
		Raw:     raw,
		Message: message,
		Code:    code,
	}
}
