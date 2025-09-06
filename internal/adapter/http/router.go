package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jefferssongalvao/go_clean_arch/internal/adapter/http/handlers"
)

func SetupRouter(studentHandler *handlers.StudentHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/students", studentHandler.GetAll)
	r.GET("/students/:id", studentHandler.GetByID)
	r.POST("/students", studentHandler.Create)
	r.PATCH("/students/:id", studentHandler.Update)
	r.DELETE("/students/:id", studentHandler.Delete)
	return r
}
