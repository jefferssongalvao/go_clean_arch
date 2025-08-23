package domain

import "errors"

// Erros do domínio
var (
	ErrInvalidStudent  = errors.New("invalid student")
	ErrStudentNotFound = errors.New("student not found")
)
