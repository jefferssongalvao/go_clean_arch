package usecase

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/domain"
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/errors"
)

type StudentService struct {
	repo domain.StudentRepository
}

func NewStudentService(r domain.StudentRepository) *StudentService {
	return &StudentService{repo: r}
}

func (s *StudentService) GetAll(name string) ([]domain.Student, error) {
	return s.repo.FindAll(name)
}

func (s *StudentService) GetByID(id uint) (*domain.Student, error) {
	return s.repo.FindByID(id)
}

func (s *StudentService) Create(student *domain.Student) error {
	if student.Name == "" {
		return errors.ErrInvalidStudent
	}
	return s.repo.Create(student)
}

func (s *StudentService) Update(student *domain.Student) error {
	if student.Name == "" {
		return errors.ErrInvalidStudent
	}
	return s.repo.Update(student)
}

func (s *StudentService) Delete(id uint) error {
	return s.repo.Delete(id)
}
