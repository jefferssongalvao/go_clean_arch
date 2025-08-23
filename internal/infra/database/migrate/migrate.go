package migrate

import (
	"log"

	"github.com/jefferssongalvao/go_clean_arch/internal/infra/database"
	"github.com/jefferssongalvao/go_clean_arch/internal/models"
)

func Run() {
	db := database.DB

	if err := db.AutoMigrate(&models.Student{}); err != nil {
		log.Panic("Failed to migrate:", err)
	}

	log.Println("Database migrated!")
}
