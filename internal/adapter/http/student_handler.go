package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	valueobjects "github.com/jefferssongalvao/go_clean_arch/internal/domain/value_objects"
)

// StudentServiceInterface define as operações esperadas pelo handler
type StudentServiceInterface interface {
	GetAll(name string) ([]entities.Student, error)
	GetByID(id uint) (*entities.Student, error)
	Create(student *entities.Student) error
	Update(student *entities.Student) error
	Delete(id uint) error
}

type StudentHandler struct {
	svc StudentServiceInterface
}

func NewStudentHandler(s StudentServiceInterface) *StudentHandler {
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
	var resp []StudentResponse
	for _, s := range students {
		resp = append(resp, StudentResponse{
			ID:    s.ID,
			Name:  s.Name,
			Email: s.Email.String(),
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
	resp := StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email.String(),
	}
	c.JSON(http.StatusOK, resp)
}

func (h *StudentHandler) Create(c *gin.Context) {
	var req StudentRequest
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
		Name:  req.Name,
		Email: email,
	}
	if err := h.svc.Create(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email.String(),
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *StudentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req StudentRequest
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
	if err := h.svc.Update(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email.String(),
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
