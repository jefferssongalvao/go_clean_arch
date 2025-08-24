package seeds

import (
	"fmt"
	"log"

	valueobjects "github.com/jefferssongalvao/go_clean_arch/internal/domain/value_objects"
	"github.com/jefferssongalvao/go_clean_arch/internal/infra/database"
	"github.com/jefferssongalvao/go_clean_arch/internal/infra/database/models"
)

func Run() {
	db := database.DB

	// Lista de usuários e estudantes relacionados
	pass1, _ := valueobjects.NewPassword("Bat@123")

	fmt.Println("Password 1: ", pass1)

	pass2, _ := valueobjects.NewPassword("super123")

	fmt.Println("Password 2: ", pass2)

	users := []models.User{
		{Username: "bruce", Password: pass1.Hash()},
		{Username: "clark", Password: pass2.Hash()},
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
