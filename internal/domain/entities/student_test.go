package entities

import (
	"testing"

	valueobjects "github.com/jefferssongalvao/go_clean_arch/internal/domain/value_objects"
)

func TestStudent_Valid(t *testing.T) {
	email, _ := valueobjects.NewEmail("user@email.com")
	s := Student{
		ID:    1,
		Name:  "Bruce Wayne",
		Email: email,
	}
	if s.Name != "Bruce Wayne" {
		t.Errorf("esperado 'Bruce Wayne', obteve '%s'", s.Name)
	}
	if !s.Email.Equals(email) {
		t.Error("emails não são iguais")
	}
}
