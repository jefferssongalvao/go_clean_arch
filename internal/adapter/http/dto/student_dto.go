package dto

type StudentRequest struct {
	Name  string      `json:"name" binding:"required"`
	Email string      `json:"email" binding:"required,email"`
	User  UserRequest `json:"user" binding:"required"`
}

type StudentResponse struct {
	ID    uint         `json:"id"`
	Name  string       `json:"name"`
	Email string       `json:"email"`
	User  UserResponse `json:"user"`
}
