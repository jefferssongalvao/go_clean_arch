package usecase

import (
	"errors"

	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
)

var (
	ErrInvalidStudent  = errors.New("invalid student")
	ErrStudentNotFound = errors.New("student not found")
)

type StudentService struct {
	repo entities.StudentRepository
}

func NewStudentService(r entities.StudentRepository) *StudentService {
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
