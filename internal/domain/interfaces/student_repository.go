package interfaces

import "github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"

type StudentRepository interface {
	FindAll(name string) ([]entities.Student, error)
	FindByID(id uint) (*entities.Student, error)
	Create(student *entities.Student) (*entities.Student, error)
	Update(student *entities.Student) (*entities.Student, error)
	Delete(id uint) error
}
