package domain

type Student struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type StudentRepository interface {
	FindAll(name string) ([]Student, error)
	FindByID(id uint) (*Student, error)
	Create(student *Student) error
	Update(student *Student) error
	Delete(id uint) error
}
