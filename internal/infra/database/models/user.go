package models

import "github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (m *User) ToEntity() *entities.User {
	return &entities.User{
		ID:       m.ID,
		Username: m.Username,
		Password: m.Password,
	}
}

func UserFromEntity(e *entities.User) *User {
	if e == nil {
		return nil
	}
	return &User{
		ID:       e.ID,
		Username: e.Username,
		Password: e.Password,
	}
}
