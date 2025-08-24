package main

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/infra/database"
	"github.com/jefferssongalvao/go_clean_arch/migrations"
	"github.com/jefferssongalvao/go_clean_arch/migrations/seeds"
)

func main() {
	database.Connect()
	migrations.Run()
	seeds.Run()
}
