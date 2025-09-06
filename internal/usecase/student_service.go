package usecase

import (
	"context"
	"errors"

	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/interfaces"
	"github.com/jefferssongalvao/go_clean_arch/internal/infra/observability"
)

var (
	ErrInvalidStudent  = errors.New("invalid student")
	ErrStudentNotFound = errors.New("student not found")
)

type IStudentService interface {
	GetAll(ctx context.Context, name string) ([]entities.Student, error)
	GetByID(ctx context.Context, id uint) (*entities.Student, error)
	Create(ctx context.Context, student *entities.Student) (*entities.Student, error)
	Update(ctx context.Context, student *entities.Student) (*entities.Student, error)
	Delete(ctx context.Context, id uint) error
}

type StudentService struct {
	repo interfaces.StudentRepository
}

func NewStudentService(r interfaces.StudentRepository) IStudentService {
	return &StudentService{repo: r}
}

func (s *StudentService) GetAll(ctx context.Context, name string) ([]entities.Student, error) {
	defer observability.StartSegment(ctx, "StudentService.GetAll")()
	return s.repo.FindAll(name)
}

func (s *StudentService) GetByID(ctx context.Context, id uint) (*entities.Student, error) {
	defer observability.StartSegment(ctx, "StudentService.GetByID")()
	return s.repo.FindByID(id)
}

func (s *StudentService) Create(ctx context.Context, student *entities.Student) (*entities.Student, error) {
	defer observability.StartSegment(ctx, "StudentService.Create")()
	if student.Name == "" {
		return nil, ErrInvalidStudent
	}
	return s.repo.Create(student)
}

func (s *StudentService) Update(ctx context.Context, student *entities.Student) (*entities.Student, error) {
	defer observability.StartSegment(ctx, "StudentService.Update")()
	if student.Name == "" {
		return nil, ErrInvalidStudent
	}
	return s.repo.Update(student)
}

func (s *StudentService) Delete(ctx context.Context, id uint) error {
	defer observability.StartSegment(ctx, "StudentService.Delete")()
	return s.repo.Delete(id)
}
