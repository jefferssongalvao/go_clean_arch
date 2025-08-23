package main

import (
	"github.com/jefferssongalvao/go_clean_arch/internal/database"
	"github.com/jefferssongalvao/go_clean_arch/internal/seeds"
)

func main() {
	database.Connect()
	seeds.Run()
}
