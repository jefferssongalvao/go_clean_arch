package models

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	valueobjects "github.com/jefferssongalvao/go_clean_arch/internal/domain/value_objects"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (m *User) ToEntity() *entities.User {
	password, err := valueobjects.NewPassword(m.Password)
	if err != nil {
		return nil
	}
	return &entities.User{
		ID:       m.ID,
		Username: m.Username,
		Password: *password,
	}
}

func UserFromEntity(e *entities.User) *User {
	if e == nil {
		return nil
	}
	return &User{
		ID:       e.ID,
		Username: e.Username,
		Password: e.Password.Hash(),
	}
}
