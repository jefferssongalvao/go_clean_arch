package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jefferssongalvao/go_clean_arch/internal/adapter/http/handlers"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func SetupRouter(
	app *newrelic.Application,
	studentHandler *handlers.StudentHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(nrgin.Middleware(app))
	students := r.Group("/students")
	{
		students.GET("", studentHandler.GetAll)
		students.GET("/:id", studentHandler.GetByID)
		students.POST("", studentHandler.Create)
		students.PATCH("/:id", studentHandler.Update)
		students.DELETE("/:id", studentHandler.Delete)
	}
	return r
}
