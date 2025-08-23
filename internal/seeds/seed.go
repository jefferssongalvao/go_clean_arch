package seeds

import (
	"log"

	"github.com/jefferssongalvao/go_clean_arch/internal/database"
	"github.com/jefferssongalvao/go_clean_arch/internal/models"
)

func Run() {
	db := database.DB

	// AutoMigrate
	if err := db.AutoMigrate(&models.Student{}); err != nil {
		log.Panic("Failed to migrate:", err)
	}

	// Seed
	students := []models.Student{
		{Name: "Bruce Wayne"},
		{Name: "Clark Kent"},
	}
	for _, s := range students {
		db.FirstOrCreate(&s, models.Student{Name: s.Name})
	}
	log.Println("Database migrated and seeded!")
}
