package valueobjects

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash string
}

func NewPassword(plain string) (*Password, error) {
	hashed, err := generateHash(plain)

	if err != nil {
		return nil, err
	}

	return &Password{hash: hashed}, nil
}

func generateHash(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (p *Password) Validate(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plain))
	return err == nil
}

func (p *Password) Hash() string {
	return p.hash
}
