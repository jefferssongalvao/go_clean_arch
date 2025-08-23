package persistence

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/domain"
	"gorm.io/gorm"
)

type studentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) domain.StudentRepository {
	return &studentRepo{db}
}

func (r *studentRepo) FindAll(name string) ([]domain.Student, error) {
	var students []domain.Student
	query := r.db
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	err := query.Find(&students).Error
	return students, err
}

func (r *studentRepo) FindByID(id uint) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepo) Create(student *domain.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepo) Update(student *domain.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepo) Delete(id uint) error {
	return r.db.Delete(&domain.Student{}, id).Error
}
