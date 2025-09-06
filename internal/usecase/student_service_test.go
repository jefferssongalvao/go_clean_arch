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
func (f *fakeRepo) Create(student *entities.Student) (*entities.Student, error) {
	f.students = append(f.students, *student)
	return student, nil
}
func (f *fakeRepo) Update(student *entities.Student) (*entities.Student, error) {
	return student, nil
}
func (f *fakeRepo) Delete(id uint) error {
	return nil
}

func TestStudentServiceCreate(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewStudentService(repo)
	student := &entities.Student{Name: "Clark Kent"}
	created, err := svc.Create(student)
	if err != nil {
		t.Errorf("esperado sucesso, obteve erro: %v", err)
	}
	if created == nil || created.Name != "Clark Kent" {
		t.Errorf("esperado estudante criado com nome 'Clark Kent', obteve %+v", created)
	}
}

func TestStudentServiceCreateInvalid(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewStudentService(repo)
	student := &entities.Student{Name: ""}
	created, err := svc.Create(student)
	if err == nil {
		t.Error("esperado erro para nome vazio")
	}
	if created != nil {
		t.Error("esperado estudante nulo para nome vazio")
	}
}
