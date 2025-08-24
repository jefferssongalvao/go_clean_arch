package main

import (
	"github.com/jefferssongalvao/go_clean_arch/infra/database"
	"github.com/jefferssongalvao/go_clean_arch/infra/database/migrate"
	"github.com/jefferssongalvao/go_clean_arch/infra/database/migrate/seeds"
)

func main() {
	database.Connect()
	migrate.Run()
	seeds.Run()
}
