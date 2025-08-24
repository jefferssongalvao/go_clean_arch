package valueobjects

import (
	"testing"
)

func TestNewEmail_Valid(t *testing.T) {
	email, err := NewEmail("user@email.com")
	if err != nil {
		t.Errorf("esperado email válido, mas retornou erro: %v", err)
	}
	if email.String() != "user@email.com" {
		t.Errorf("esperado 'user@email.com', obteve '%s'", email.String())
	}
}

func TestNewEmail_Invalid(t *testing.T) {
	_, err := NewEmail("invalido-sem-arroba")
	if err == nil {
		t.Error("esperado erro para email inválido, mas não retornou erro")
	}
}
