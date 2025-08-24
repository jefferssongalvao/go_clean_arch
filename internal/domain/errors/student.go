package errors

import "errors"

var (
	ErrInvalidStudent  = errors.New("invalid student")
	ErrStudentNotFound = errors.New("student not found")
)
