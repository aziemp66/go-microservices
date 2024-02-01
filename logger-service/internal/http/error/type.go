package http_error

import "fmt"

type Error struct {
	Raw     error
	Message string
	Code    int
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error %d : %s", e.Code, e.Message)
}
