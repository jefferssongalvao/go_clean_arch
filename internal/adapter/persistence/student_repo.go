package persistence

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	"gorm.io/gorm"
)

type studentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) entities.StudentRepository {
	return &studentRepo{db}
}

func (r *studentRepo) FindAll(name string) ([]entities.Student, error) {
	var students []entities.Student
	query := r.db
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	err := query.Find(&students).Error
	return students, err
}

func (r *studentRepo) FindByID(id uint) (*entities.Student, error) {
	var student entities.Student
	if err := r.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepo) Create(student *entities.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepo) Update(student *entities.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepo) Delete(id uint) error {
	return r.db.Delete(&entities.Student{}, id).Error
}
