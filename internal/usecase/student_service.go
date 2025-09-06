package usecase

import (
	"errors"

	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/interfaces"
)

var (
	ErrInvalidStudent  = errors.New("invalid student")
	ErrStudentNotFound = errors.New("student not found")
)

type IStudentService interface {
	GetAll(name string) ([]entities.Student, error)
	GetByID(id uint) (*entities.Student, error)
	Create(student *entities.Student) (*entities.Student, error)
	Update(student *entities.Student) (*entities.Student, error)
	Delete(id uint) error
}

type StudentService struct {
	repo interfaces.StudentRepository
}

func NewStudentService(r interfaces.StudentRepository) IStudentService {
	return &StudentService{repo: r}
}

func (s *StudentService) GetAll(name string) ([]entities.Student, error) {
	return s.repo.FindAll(name)
}

func (s *StudentService) GetByID(id uint) (*entities.Student, error) {
	return s.repo.FindByID(id)
}

func (s *StudentService) Create(student *entities.Student) (*entities.Student, error) {
	if student.Name == "" {
		return nil, ErrInvalidStudent
	}
	return s.repo.Create(student)
}

func (s *StudentService) Update(student *entities.Student) (*entities.Student, error) {
	if student.Name == "" {
		return nil, ErrInvalidStudent
	}
	return s.repo.Update(student)
}

func (s *StudentService) Delete(id uint) error {
	return s.repo.Delete(id)
}
