package seeds

import (
	"log"

	"github.com/jefferssongalvao/go_clean_arch/internal/infra/database"
	"github.com/jefferssongalvao/go_clean_arch/internal/infra/database/models"
)

func Run() {
	db := database.DB

	// Lista de usuários e estudantes relacionados
	users := []models.User{
		{Username: "bruce", Password: "bat123"},
		{Username: "clark", Password: "super123"},
	}

	students := []models.Student{
		{Name: "Bruce Wayne", Email: "bruce@gmail.com"},
		{Name: "Clark Kent", Email: "clark@gmail.com"},
	}

	for i, user := range users {
		// Cria o usuário
		if err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error; err != nil {
			log.Fatalf("Erro ao criar usuário: %v", err)
		}
		// Associa o estudante ao usuário criado
		students[i].UserID = user.ID
		if err := db.FirstOrCreate(&students[i], models.Student{Name: students[i].Name}).Error; err != nil {
			log.Fatalf("Erro ao criar estudante: %v", err)
		}
	}

	log.Println("Database seeded!")
}
