package seeds

import (
	"log"

	"github.com/jefferssongalvao/go_clean_arch/infra/database"
	"github.com/jefferssongalvao/go_clean_arch/infra/database/models"
)

func Run() {
	db := database.DB

	students := []models.Student{
		{Name: "Bruce Wayne", Email: "bruce@gmail.com"},
		{Name: "Clark Kent", Email: "clark@gmail.com"},
	}

	for _, s := range students {
		db.FirstOrCreate(&s, models.Student{Name: s.Name})
	}

	log.Println("Database seeded!")
}
