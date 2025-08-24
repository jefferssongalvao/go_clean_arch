package persistence

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	"github.com/jefferssongalvao/go_clean_arch/internal/infra/database/models"
	"gorm.io/gorm"
)

type studentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) entities.StudentRepository {
	return &studentRepo{db}
}

func (r *studentRepo) FindAll(name string) ([]entities.Student, error) {
	var students []models.Student
	query := r.db.Preload("User")
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	err := query.Find(&students).Error

	var result []entities.Student
	for _, student := range students {
		result = append(result, *student.ToEntity())
	}

	return result, err
}

func (r *studentRepo) FindByID(id uint) (*entities.Student, error) {
	var student models.Student
	if err := r.db.Preload("User").First(&student, id).Error; err != nil {
		return nil, err
	}

	return (&student).ToEntity(), nil
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
