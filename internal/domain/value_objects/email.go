package valueobjects

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
)

// Email representa um Value Object de endereço de email.
type Email struct {
	value string
}

// regex básica para validar email (não cobre todos os casos possíveis, mas é prática para uso geral)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// NewEmail cria um novo Email validando o valor.
func NewEmail(address string) (Email, error) {
	address = strings.TrimSpace(address)
	address = strings.ToLower(address)

	if !emailRegex.MatchString(address) {
		return Email{}, ErrInvalidEmail
	}

	return Email{value: address}, nil
}

// String retorna o valor do Email como string.
func (e Email) String() string {
	return e.value
}

// Equals compara dois Value Objects de Email.
func (e Email) Equals(other Email) bool {
	return e.value == other.value
}
