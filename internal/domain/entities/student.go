package entities

import valueobjects "github.com/jefferssongalvao/go_clean_arch/internal/domain/value_objects"

type Student struct {
	ID     uint               `json:"id"`
	Name   string             `json:"name"`
	Email  valueobjects.Email `json:"email"`
	UserID uint               `json:"user_id"`
	User   *User              `json:"user"`
}

type StudentRepository interface {
	FindAll(name string) ([]Student, error)
	FindByID(id uint) (*Student, error)
	Create(student *Student) error
	Update(student *Student) error
	Delete(id uint) error
}

func EmailFromString(s string) valueobjects.Email {
	email, _ := valueobjects.NewEmail(s)
	return email
}
