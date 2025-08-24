package usecase

import (
	"errors"
	"testing"

	entities "github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
)

type fakeRepo struct {
	students []entities.Student
}

func (f *fakeRepo) FindAll(name string) ([]entities.Student, error) {
	return f.students, nil
}
func (f *fakeRepo) FindByID(id uint) (*entities.Student, error) {
	for _, s := range f.students {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, errors.New("not found")
}
func (f *fakeRepo) Create(student *entities.Student) error {
	f.students = append(f.students, *student)
	return nil
}
func (f *fakeRepo) Update(student *entities.Student) error {
	return nil
}
func (f *fakeRepo) Delete(id uint) error {
	return nil
}

func TestStudentService_Create(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewStudentService(repo)
	student := &entities.Student{Name: "Clark Kent"}
	err := svc.Create(student)
	if err != nil {
		t.Errorf("esperado sucesso, obteve erro: %v", err)
	}
}

func TestStudentService_Create_Invalid(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewStudentService(repo)
	student := &entities.Student{Name: ""}
	err := svc.Create(student)
	if err == nil {
		t.Error("esperado erro para nome vazio")
	}
}
