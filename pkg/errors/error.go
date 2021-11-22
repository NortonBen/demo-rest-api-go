package errors

import "encoding/json"

type Error struct {
	Code      int
	Message   string
	ShortCode string
}

func (e *Error) Error() string {
	rs, _ := json.Marshal(e)
	return string(rs)
}

func BadRequest(message string, shortCode string) error {
	return &Error{
		Code:      400,
		Message:   message,
		ShortCode: shortCode,
	}
}

func NotFound(message string) error {
	return &Error{
		Code:      404,
		Message:   message,
		ShortCode: "not-found",
	}
}
