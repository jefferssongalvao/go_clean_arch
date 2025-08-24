package models

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Email  string `gorm:"uniqueIndex;not null"`
	UserID uint   `gorm:"not null;unique"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:ID"`
}

func (m *Student) ToEntity() *entities.Student {
	return &entities.Student{
		ID:     m.ID,
		Name:   m.Name,
		Email:  entities.EmailFromString(m.Email), // crie essa função utilitária
		UserID: m.UserID,
		User:   m.User.ToEntity(),
	}
}

func StudentFromEntity(e *entities.Student) *Student {
	return &Student{
		Model:  gorm.Model{ID: e.ID},
		Name:   e.Name,
		Email:  e.Email.String(),
		UserID: e.UserID,
		// Não preencher User ao criar, apenas UserID
	}
}
