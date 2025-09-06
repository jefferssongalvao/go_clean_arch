package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jefferssongalvao/go_clean_arch/internal/adapter/http/dto"
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	valueobjects "github.com/jefferssongalvao/go_clean_arch/internal/domain/value_objects"
	"github.com/jefferssongalvao/go_clean_arch/internal/usecase"
)

type StudentHandler struct {
	svc usecase.IStudentService
}

func NewStudentHandler(s usecase.IStudentService) *StudentHandler {
	return &StudentHandler{svc: s}
}

func (h *StudentHandler) GetAll(c *gin.Context) {
	name := c.Query("name")
	students, err := h.svc.GetAll(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(students) == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	// Mapear para DTO de resposta
	var resp []dto.StudentResponse
	for _, s := range students {
		resp = append(resp, dto.StudentResponse{
			ID:    s.ID,
			Name:  s.Name,
			Email: s.Email.String(),
			User: dto.UserResponse{
				ID:       s.User.ID,
				Username: s.User.Username,
			},
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	student, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	resp := dto.StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email.String(),
		User: dto.UserResponse{
			ID:       student.User.ID,
			Username: student.User.Username,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func (h *StudentHandler) Create(c *gin.Context) {
	var req dto.StudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email, err := valueobjects.NewEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}
	pass, err := valueobjects.NewPassword(req.User.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	student := entities.Student{
		Name:  req.Name,
		Email: email,
		User: &entities.User{
			Username: req.User.Username,
			Password: *pass,
		},
	}

	respCreate, err := h.svc.Create(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := dto.StudentResponse{
		ID:    respCreate.ID,
		Name:  respCreate.Name,
		Email: respCreate.Email.String(),
		User: dto.UserResponse{
			ID:       respCreate.User.ID,
			Username: respCreate.User.Username,
		},
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *StudentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.StudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email, err := valueobjects.NewEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}
	student := entities.Student{
		ID:    uint(id),
		Name:  req.Name,
		Email: email,
	}
	respUpdate, err := h.svc.Update(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := dto.StudentResponse{
		ID:    respUpdate.ID,
		Name:  respUpdate.Name,
		Email: respUpdate.Email.String(),
		User: dto.UserResponse{
			ID:       respUpdate.User.ID,
			Username: respUpdate.User.Username,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
