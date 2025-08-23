package main

import (
	"log"

	"github.com/jefferssongalvao/go_clean_arch/config"
	"github.com/jefferssongalvao/go_clean_arch/internal/adapter/http"
	"github.com/jefferssongalvao/go_clean_arch/internal/adapter/persistence"
	"github.com/jefferssongalvao/go_clean_arch/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.GetDSN()))
	if err != nil {
		log.Fatal(err)
	}

	repo := persistence.NewStudentRepo(db)
	svc := usecase.NewStudentService(repo)
	handler := http.NewStudentHandler(svc)

	r := http.SetupRouter(handler)
	r.Run(":8080")
}
