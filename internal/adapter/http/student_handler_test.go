package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	entities "github.com/jefferssongalvao/go_clean_arch/internal/domain/entities"
	valueobjects "github.com/jefferssongalvao/go_clean_arch/internal/domain/value_objects"
)

type fakeService struct{}

func (f *fakeService) GetAll(name string) ([]entities.Student, error) {
	return []entities.Student{}, nil
}
func (f *fakeService) GetByID(id uint) (*entities.Student, error) {
	return &entities.Student{ID: id, Name: "Test", Email: valueobjects.Email{}}, nil
}
func (f *fakeService) Create(student *entities.Student) error { return nil }
func (f *fakeService) Update(student *entities.Student) error { return nil }
func (f *fakeService) Delete(id uint) error                   { return nil }

func TestStudentHandler_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewStudentHandler(&fakeService{})
	r := gin.Default()
	r.GET("/students", h.GetAll)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/students", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent && w.Code != http.StatusOK {
		t.Errorf("esperado 204 ou 200, obteve %d", w.Code)
	}
}

func TestStudentHandler_Create_Invalid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewStudentHandler(&fakeService{})
	r := gin.Default()
	r.POST("/students", h.Create)
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{"name": "", "email": "invalido"})
	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("esperado 400, obteve %d", w.Code)
	}
}
