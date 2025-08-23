package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(studentHandler *StudentHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/students", studentHandler.GetAll)
	r.GET("/students/:id", studentHandler.GetByID)
	r.POST("/students", studentHandler.Create)
	r.PATCH("/students/:id", studentHandler.Update)
	r.DELETE("/students/:id", studentHandler.Delete)
	return r
}
