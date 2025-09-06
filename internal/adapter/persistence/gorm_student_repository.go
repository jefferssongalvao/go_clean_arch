package persistence

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/interfaces"
	"github.com/jefferssongalvao/go_clean_arch/internal/infra/database/models"
	"gorm.io/gorm"
)

type GormStudentRepository struct {
	db *gorm.DB
}

func NewGormStudentRepository(db *gorm.DB) interfaces.StudentRepository {
	return &GormStudentRepository{db}
}

func (r *GormStudentRepository) FindAll(name string) ([]entities.Student, error) {
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

func (r *GormStudentRepository) FindByID(id uint) (*entities.Student, error) {
	var student models.Student
	if err := r.db.Preload("User").First(&student, id).Error; err != nil {
		return nil, err
	}

	return (&student).ToEntity(), nil
}

func (r *GormStudentRepository) Create(student *entities.Student) (*entities.Student, error) {
	var createdStudent *entities.Student
	err := r.db.Transaction(func(tx *gorm.DB) error {
		userModel := models.UserFromEntity(student.User)
		if err := tx.Create(userModel).Error; err != nil {
			return err
		}

		student.UserID = userModel.ID
		studentModel := models.StudentFromEntity(student)

		if err := tx.Create(studentModel).Error; err != nil {
			return err
		}

		// Buscar estudante criado com preload do usuário
		var dbStudent models.Student
		if err := tx.Preload("User").First(&dbStudent, studentModel.ID).Error; err != nil {
			return err
		}
		createdStudent = dbStudent.ToEntity()
		return nil
	})
	return createdStudent, err
}

func (r *GormStudentRepository) Update(student *entities.Student) (*entities.Student, error) {
	studentModel := models.StudentFromEntity(student)
	if err := r.db.Save(studentModel).Error; err != nil {
		return nil, err
	}
	// Buscar estudante atualizado com preload do usuário
	var dbStudent models.Student
	if err := r.db.Preload("User").First(&dbStudent, studentModel.ID).Error; err != nil {
		return nil, err
	}
	return dbStudent.ToEntity(), nil
}

func (r *GormStudentRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Student{}, id).Error
}
